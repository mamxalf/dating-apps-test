package constant

import "fmt"

var (
	// ErrWrongPassword is
	ErrWrongPassword = fmt.Errorf("invalid password")
	// ErrCannotSendOTP is
	ErrCannotSendOTP = fmt.Errorf("cannot send otp code, please try again later")
	// ErrOnHoldOTPInput is
	ErrOnHoldOTPInput = fmt.Errorf("your account is temporarily suspended")
	// ErrWrongOTPCode is
	ErrWrongOTPCode = fmt.Errorf("invalid otp code")
	// ErrExpiredOTP is
	ErrExpiredOTP = fmt.Errorf("otp code is expired")
	// ErrTokenNotFound is
	ErrTokenNotFound = fmt.Errorf("invalid or expired jwt")
	// ErrUserCategory is
	ErrUserCategory = fmt.Errorf("only personal account is allowed")
	// ErrMSCreateUser is
	ErrMSCreateUser = fmt.Errorf("failed to create user")
	// ErrMSUpdateUser is
	ErrMSUpdateUser = fmt.Errorf("failed to update user")
	// ErrInvalidReferralCode is
	ErrInvalidReferralCode = fmt.Errorf("invalid referral code")
	// ErrInvalidAuthorization is
	ErrInvalidAuthorization = fmt.Errorf("invalid authorization")
	// ErrUserAlreadySuspended is
	ErrUserAlreadySuspended = fmt.Errorf("user already suspended")
	// ErrUserAlreadyUnsuspended is
	ErrUserAlreadyUnsuspended = fmt.Errorf("user already unsuspended")
	// ErrUserSuspended is
	ErrUserSuspended = fmt.Errorf("user suspended")
	// ErrPinAlreadyExists is
	ErrPinAlreadyExists = fmt.Errorf("pin has already been set")
	// ErrSessionNotExists is
	ErrSessionNotExists = fmt.Errorf("session doesn't exists")
	// ErrInvalidSignedKey is
	ErrInvalidSignedKey = fmt.Errorf("invalid signed key")
	// ErrWrongPIN is
	ErrWrongPIN = fmt.Errorf("invalid pin")
	// ErrWrongPIN is
	ErrWrongOldPIN = fmt.Errorf("invalid old pin")
	// ErrEmailHasVerified is
	ErrEmailHasVerified = fmt.Errorf("your email has been verified")
	// ErrOnHoldSendEmail is
	ErrOnHoldSendEmail = fmt.Errorf("cannot send email, your account is temporarily suspended")
	// ErrCannotSendEmail is
	ErrCannotSendEmail = fmt.Errorf("cannot send email, please try again later")
	// ErrPhoneNumberRegistered is
	ErrPhoneNumberRegistered = fmt.Errorf("phone number is already registered")
	// ErrLoginPinNotAllowed is
	ErrLoginPinNotAllowed = fmt.Errorf("login with pin is not allowed")
	// ErrOnHoldLoginPIN is
	ErrOnHoldLoginPIN = fmt.Errorf("too many failures, login with pin is temporarily suspended")
	// ErrCannotSendOTP is
	ErrWaitToSendOTP = "wait %d second to request again"
	// ErrEmailIsNotEqual is
	ErrEmailIsNotEqual = fmt.Errorf("email is not equal to existing email")
	// ErrInvalidEmailToken is
	ErrInvalidEmailToken = fmt.Errorf("invalid email token")
	// ErrInvalidPin is
	ErrInvalidPinIsSequential = fmt.Errorf("pin can't be sequential")
	// ErrInvalidPinIsRepeated is
	ErrInvalidPinIsRepeated = fmt.Errorf("pin can't repeat like `11`32")
	// ErrInvalidEmailOtp is
	ErrInvalidEmailOtp = fmt.Errorf("invalid email otp")
	// ErrPasswordNotSetup is
	ErrPasswordNotSetup = fmt.Errorf("password is not setup")
	// ErrEmailNotVerified is
	ErrEmailNotVerified = fmt.Errorf("email not verified")
	// ErrExpiredAuthorization is
	ErrExpiredAuthorization = fmt.Errorf("authorization token is expired")
	// ErrAuthorizationNotProvided is
	ErrAuthorizationNotProvided = fmt.Errorf("authorization token is not provided")
	// ErrCantUpdateClientKeyAndSecretForPersonalUser is
	ErrCantUpdateClientKeyAndSecretForPersonalUser = fmt.Errorf("can't update client key and client secret for personal user")
	// ErrInvalidOldPassword is
	ErrInvalidOldPassword = fmt.Errorf("invalid old password")
	// ErrChangePassword is
	ErrChangePassword = fmt.Errorf("change password failed")
	// ErrPasswordIsSetup is
	ErrPasswordIsSetup = fmt.Errorf("password is not setup")
	// ErrConvertMsisdn is
	ErrMsisdnDefault       = fmt.Errorf("cannot create default msisdn")
	ErrCannotGetUserDetail = fmt.Errorf("cannot get user detail")
	ErrInternalServerError = fmt.Errorf("internal server error")
	ErrMinPasswordLength   = fmt.Errorf("minimal password length is 8 character")
)
