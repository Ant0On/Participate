package models

import (
	"fmt"

	"gorm.io/gorm"
)

type Country struct {
	gorm.Model
	CountryName string
	Towns       []Town
}

var CountryList = []Country{
	{CountryName: "China"},
	{CountryName: "India"},
	{CountryName: "United States"},
	{CountryName: "Indonesia"},
	{CountryName: "Pakistan"},
	{CountryName: "Brazil"},
	{CountryName: "Nigeria"},
	{CountryName: "Bangladesh"},
	{CountryName: "Russia"},
	{CountryName: "Mexico"},
	{CountryName: "Japan"},
	{CountryName: "Ethiopia"},
	{CountryName: "Philippines"},
	{CountryName: "Egypt"},
	{CountryName: "Vietnam"},
	{CountryName: "Congo"},
	{CountryName: "Turkey"},
	{CountryName: "Iran"},
	{CountryName: "Germany"},
	{CountryName: "France"},
	{CountryName: "United Kingdom"},
	{CountryName: "Thailand"},
	{CountryName: "Italy"},
	{CountryName: "Tanzania"},
	{CountryName: "South Africa"},
	{CountryName: "Myanmar"},
	{CountryName: "South Korea"},
	{CountryName: "Kenya"},
	{CountryName: "Colombia"},
	{CountryName: "Spain"},
	{CountryName: "Argentina"},
	{CountryName: "Algeria"},
	{CountryName: "Sudan"},
	{CountryName: "Ukraine"},
	{CountryName: "Uganda"},
	{CountryName: "Canada"},
	{CountryName: "Iraq"},
	{CountryName: "Morocco"},
	{CountryName: "Afghanistan"},
	{CountryName: "Poland"},
	{CountryName: "Malaysia"},
	{CountryName: "Venezuela"},
	{CountryName: "Peru"},
	{CountryName: "Uzbekistan"},
	{CountryName: "Saudi Arabia"},
	{CountryName: "Mozambique"},
	{CountryName: "Ghana"},
	{CountryName: "Yemen"},
	{CountryName: "Nepal"},
	{CountryName: "Venezuela"},
	{CountryName: "Madagascar"},
	{CountryName: "Cameroon"},
	{CountryName: "Ivory Coast"},
	{CountryName: "Niger"},
	{CountryName: "Mali"},
	{CountryName: "Malawi"},
	{CountryName: "Kazakhstan"},
	{CountryName: "Ecuador"},
	{CountryName: "Guatemala"},
	{CountryName: "Netherlands"},
	{CountryName: "Syria"},
	{CountryName: "Cambodia"},
	{CountryName: "Senegal"},
	{CountryName: "Chad"},
	{CountryName: "Somalia"},
	{CountryName: "Zimbabwe"},
	{CountryName: "Rwanda"},
	{CountryName: "Benin"},
	{CountryName: "Burundi"},
	{CountryName: "Tunisia"},
	{CountryName: "Bolivia"},
	{CountryName: "Belgium"},
	{CountryName: "Czech Republic"},
	{CountryName: "Greece"},
	{CountryName: "Portugal"},
	{CountryName: "Sweden"},
	{CountryName: "Azerbaijan"},
	{CountryName: "Hungary"},
	{CountryName: "Belarus"},
	{CountryName: "Austria"},
	{CountryName: "Switzerland"},
	{CountryName: "Bolivia"},
	{CountryName: "Jordan"},
	{CountryName: "Serbia"},
	{CountryName: "Libya"},
	{CountryName: "United Arab Emirates"},
	{CountryName: "Honduras"},
	{CountryName: "Lebanon"},
	{CountryName: "Burma"},
	{CountryName: "Lebanon"},
	{CountryName: "Singapore"},
	{CountryName: "Belarus"},
	{CountryName: "Sri Lanka"},
	{CountryName: "Luxembourg"},
	{CountryName: "Kyrgyzstan"},
	{CountryName: "Turkmenistan"},
	{CountryName: "Sierra Leone"},
	{CountryName: "Tajikistan"},
	{CountryName: "Cyprus"},
	{CountryName: "Palestine"},
}

func (c *Country) save() error {
	if err := DB.Create(&c).Error; err != nil {
		return fmt.Errorf("failed to create country: %w", err)
	}
	return nil
}

func AddCountries() error {
	for _, country := range CountryList {
		if err := country.save(); err != nil {
			return fmt.Errorf("failed to add countries: %w", err)
		}
	}
	return nil
}

func GetAllCountries() ([]Country, error) {
	var countries []Country

	if err := DB.Order("country_name").Find(&countries).Error; err != nil {
		return nil, fmt.Errorf("failed to retrieve countries: %w", err)
	}
	return countries, nil
}
