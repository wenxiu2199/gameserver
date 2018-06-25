package poke

import (
	"errors"
	"fmt"
	"math/rand"

	"server/manager"

	"github.com/golang/protobuf/proto"
	"github.com/sirupsen/logrus"
	"github.com/wenxiu2199/gameserver/src/server/gameproto/cmsg"
	"github.com/wenxiu2199/gameserver/src/server/gameproto/gameconf"
	"github.com/wenxiu2199/gameserver/src/server/gameproto/gamedef"
)

var ErrNoSkill = errors.New("general has no skills")

type GameGeneral struct {
	*Player

	gamedef.GameGeneral
	buffs map[gameconf.SkillEffectTyp]*gamedef.Buff
}

const levelMax = 6

func newGameGeneral(general *gamedef.General, player *Player) (*GameGeneral, error) {
	gg := &GameGeneral{}
	gg.GameGeneral = gamedef.GameGeneral{}
	gg.GameGeneral.GeneralID = general.GeneralID
	gg.Player = player
	gg.buffs = make(map[gameconf.SkillEffectTyp]*gamedef.Buff)
	gg.calculateBase(general)
	gg.CurHP = gg.BaseHp
	gg.Skills = general.Skills
	gg.Level = general.Level
	gg.GameGeneral.UserID = player.ID()

	return gg, nil
}

func (p *GameGeneral) calculateBase(general *gamedef.General) error {
	conf, exist := p.getConfig().GetConfig().GetGeneralConfByGeneralID(general.GeneralID)
	if !exist {
		return fmt.Errorf("calculateBase: GetGeneralConfByGeneralID %v", general.GeneralID)
	}

	p.BaseHp = (conf.Hp*2+general.Individual.Hp+general.Effort.Hp/4)*int32(general.Level)/100 + int32(general.Level) + 10
	p.BaseAttack = (conf.Atk*2+general.Individual.Attack+general.Effort.Attack/4)*int32(general.Level)/100 + 5
	p.BaseDefense = (conf.Def*2+general.Individual.Defense+general.Effort.Defense/4)*int32(general.Level)/100 + 5
	p.BaseSpAttack = (conf.Satk*2+general.Individual.SpAttack+general.Effort.SpAttack/4)*int32(general.Level)/100 + 5
	p.BaseSpDefense = (conf.Sdef*2+general.Individual.SpDefense+general.Effort.SpDefense/4)*int32(general.Level)/100 + 5
	p.BaseSpeed = (conf.Spd*2+general.Individual.Speed+general.Effort.Speed/4)*int32(general.Level)/100 + 5

	return nil
}

func (p *GameGeneral) checkSkill(skillID uint32) bool {
	for _, v := range p.GameGeneral.Skills {
		if v == skillID {
			return true
		}
	}
	return false
}

func (p *GameGeneral) getCfg() *manager.ConfManager {
	return p.Player.GamePoke.ConfManager
}

func (p *GameGeneral) useSkill(skillID uint32, op *GameGeneral) error {
	ok := p.checkSkill(skillID)
	if !ok {
		return ErrNoSkill
	}

	skillConf, exist := p.getCfg().GetConfig().GetSkillConfBySkillID(skillID)
	if !exist {
		return ErrNoSkill
	}

	p.notifyMessage(&cmsg.CNotifyUseSkill{
		UserID:  p.Player.ID(),
		SkillID: skillID,
	})
	switch skillConf.SkillAttackType {
	case gameconf.SkillAttackTyp_SATChange:
		p.effect(skillConf, op)
	default:
		p.damage(skillConf, op)
	}
	return nil
}

func (p *GameGeneral) getAttack() int32 {
	buff, exist := p.buffs[gameconf.SkillEffectTyp_SETAttack]
	if exist {
		return p.BaseAttack*buff.Level*50/100 + p.BaseAttack
	}
	return p.BaseAttack
}

func (p *GameGeneral) getSpAttack() int32 {

	return p.BaseSpAttack
}

func (p *GameGeneral) notifyMessage(message proto.Message) {
	p.Player.GamePoke.notifyMessage(message)
}

func (p *GameGeneral) getDefense() int32 {

	return p.BaseDefense
}

func (p *GameGeneral) getSpDefense() int32 {

	return p.BaseSpDefense
}

func (p *GameGeneral) getSpeed() int32 {

	return p.BaseSpeed
}

func (p *GameGeneral) die() bool {
	if p.CurHP <= 0 {
		return true
	}
	return false
}

func (p *GameGeneral) getStatus(visible bool) *gamedef.GameGeneral {
	cp := p.GameGeneral

	if visible {
		for _, v := range p.Skills {
			cp.Skills = append(cp.Skills, v)
		}
	}

	return &cp
}

func (p *GameGeneral) damage(cfg *gameconf.SkillConfDefine, op *GameGeneral) error {
	var attack int32 = 0
	var defense int32 = 0
	switch cfg.SkillAttackType {
	case gameconf.SkillAttackTyp_SATAttack:
		attack = p.getAttack()
		defense = op.getDefense()
	case gameconf.SkillAttackTyp_SATSpecial:
		attack = p.getSpAttack()
		defense = op.getSpDefense()
	}

	damage := (float32(p.Level)*0.4+2)*float32(cfg.Power)*float32(attack)/float32(defense)/50 + 2

	op.CurHP -= int32(damage)

	p.notifyMessage(&cmsg.CNotifyGeneralStatus{
		GameGeneral: op.getStatus(true),
	})

	if op.die() {
		p.Player.GamePoke.winner = p.Player.ID()
		// todo 游戏结束
	}

	return nil
}

func (p *GameGeneral) getBuff(conf *gameconf.SkillEffectConfDefine, last int32, power int32) {
	buff, exist := p.buffs[conf.SkillEffectType]
	var level int32 = 1
	if exist {
		if last != -1 {
			last = buff.Last + last
		}
		if conf.LevelUp == 1 {
			level := level + power
			if level > levelMax {
				level = level
			}
		}
	}

	p.buffs[conf.SkillEffectType] = &gamedef.Buff{
		BuffType: conf.SkillEffectType,
		Last:     last,
		Level:    level,
	}

}

func (p *GameGeneral) effect(cfg *gameconf.SkillConfDefine, op *GameGeneral) error {
	for _, v := range cfg.Effect {
		seed := rand.Intn(100)
		if seed > int(v.Chance) {
			continue
		}
		effect, exist := p.getCfg().GetConfig().GetEffectConfByEffectID(v.Id)
		if !exist {
			logrus.WithFields(logrus.Fields{
				"effectid": v.Id,
			}).Error("GameGeneral.effet")
			continue
		}

		if v.Object == gameconf.SkillEffectObjectTyp_SEOTSelf {
			p.getBuff(effect, v.Min, v.Power)
		} else {
			op.getBuff(effect, v.Min, v.Power)
		}
	}
	return nil
}
