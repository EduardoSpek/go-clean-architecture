package validations

type InfoValidationRepository interface {
	UserWithInfo(id_user string) error
}

type InfoValidation struct {
	InfoRepository InfoValidationRepository
	UserRepository UserValidationRepository
}

func NewInfoValidation(repository InfoValidationRepository, user_repository UserValidationRepository) *InfoValidation {
	return &InfoValidation{InfoRepository: repository, UserRepository: user_repository}
}

func (v *InfoValidation) UserWithInfo(id_user string) error {
	err := v.InfoRepository.UserWithInfo(id_user)
	if err != nil {
		return err
	}
	return nil
}