package services

import (
	"example.com/m/wallet"
	dto "example.com/m/wallet/dto"
	"gorm.io/gorm"
)

type ProfileHandler struct {
	db *gorm.DB
}

func NewProfileService(db *gorm.DB) *ProfileHandler {
	return &ProfileHandler{
		db: db,
	}
}

func (p *ProfileHandler) CreateUser(dto dto.UserDto) (wallet.User, error) {
	user := wallet.User{
		IdentityType:   dto.IdentityType,
		Type:           dto.Type,
		Sub:            dto.Sub,
		IdentityNumber: dto.IdentityNumber,
		PhoneNumber:    dto.PhoneNumber,
		Email:          dto.Email,
		Address:        dto.Address,
	}
	if err := p.db.Create(&user).Error; err != nil {
		return wallet.User{}, err
	}
	return user, nil
}

func (p *ProfileHandler) FilterUsers(filters wallet.User) ([]wallet.User, error) {
	var users []wallet.User
	result := p.db.Where(filters).Find(&users)
	if result.RowsAffected == 0 {
		return users, nil
	}
	if result.Error != nil {
		return users, result.Error
	}
	return users, nil
}

func (p *ProfileHandler) CreateProfile(dto dto.ProfileDto, user wallet.User) (wallet.Profile, error) {
	profile := wallet.Profile{
		FirstName:   dto.FirstName,
		LastName:    dto.LastName,
		Gender:      dto.Gender,
		DateOfBirth: dto.DateOfBirth,
		County:      dto.County,
		SubCounty:   dto.SubCounty,
		UserID:      user.ID,
		User:        user,
	}
	if err := p.db.Create(&profile).Error; err != nil {
		return wallet.Profile{}, err
	}
	return profile, nil
}

func (p *ProfileHandler) FilterProfiles(filters wallet.Profile) ([]wallet.Profile, error) {
	var profiles []wallet.Profile
	result := p.db.Where(filters).Find(&profiles)
	if result.RowsAffected == 0 {
		return profiles, nil
	}
	if result.Error != nil {
		return profiles, result.Error
	}
	return profiles, nil
}

func (p *ProfileHandler) CreateMerchant(dto dto.MerchantDto) (wallet.Merchant, error) {
	merchant := wallet.Merchant{
		BusinessName:       dto.BusinessName,
		Email:              dto.Email,
		RegistrationNumber: dto.RegistrationNumber,
		Description:        dto.Description,
		Sub:                dto.Sub,
	}
	if err := p.db.Create(&merchant).Error; err != nil {
		return wallet.Merchant{}, err
	}
	return merchant, nil
}

func (p *ProfileHandler) FilterMerchants(filters wallet.Merchant) ([]wallet.Merchant, error) {
	var merchants []wallet.Merchant
	result := p.db.Where(filters).Find(&merchants)
	if result.RowsAffected == 0 {
		return merchants, nil
	}
	if result.Error != nil {
		return merchants, result.Error
	}
	return merchants, nil
}
