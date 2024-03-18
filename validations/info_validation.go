package validations

type InfoValidationRepository interface {
	InfoExists(id_user string) error
}

type InfoValidation struct {
	InfoRepository InfoValidationRepository
	UserRepository UserValidationRepository
}

func NewInfoValidation(repository InfoValidationRepository, user_repository UserValidationRepository) *InfoValidation {
	return &InfoValidation{InfoRepository: repository, UserRepository: user_repository}
}

func (v *InfoValidation) InfoExists(id_user string) error {
	err := v.InfoRepository.InfoExists(id_user)
	if err != nil {
		return err
	}
	return nil
}
