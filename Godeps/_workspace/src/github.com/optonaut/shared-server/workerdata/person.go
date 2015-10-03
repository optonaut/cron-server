package workerdata

import (
	"time"

	"github.com/optonaut/shared-server/utils"
)

var PersonRequestInviteTopic string

type PersonRequestInvitePayload struct {
	Email              string     `json:"email"`
	InviteActivationID utils.UUID `json:"invite_activation_id"`
}

var PersonForgotPasswordTopic string

type PersonForgotPasswordPayload struct {
	Email string `json:"email"`
}

var CronHourlyTopic string

type CronHourlyPayload struct {
	Time time.Time `json:"time"`
}

var PersonInviteActivationActivateTopic string

type PersonInviteActivationActivatePayload struct {
	Email string     `json:"email"`
	ID    utils.UUID `json:"id"`
}

func init() {
	PersonRequestInviteTopic = "person_forgot_password"
	PersonForgotPasswordTopic = "person_request_invite"
	CronHourlyTopic = "cron_hourly"
	PersonInviteActivationActivateTopic = "person_invite_activation_activate"
}
