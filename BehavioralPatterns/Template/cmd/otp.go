package cmd

type IOtp interface {
	genRandomOTP(int) string
	saveOTPCache(string)
	getMessage(string) string
	sendNotification(string) error
}

type Otp struct {
	IOtp IOtp
}

func (o *Otp) GenAndSendOTP(otpLength int) error {
	otp := o.IOtp.genRandomOTP(otpLength)
	o.IOtp.saveOTPCache(otp)
	message := o.IOtp.getMessage(otp)
	err := o.IOtp.sendNotification(message)
	if err != nil {
		return err
	}
	return nil
}
