package models

import (
	"fmt"

	"gorm.io/gorm"
)

type Country struct {
	gorm.Model
	Name  string
	Towns []Town
}

var CountryList = []Country{
	{Name: "China"},
	{Name: "India"},
	{Name: "United States"},
	{Name: "Indonesia"},
	{Name: "Pakistan"},
	{Name: "Brazil"},
	{Name: "Nigeria"},
	{Name: "Bangladesh"},
	{Name: "Russia"},
	{Name: "Mexico"},
	{Name: "Japan"},
	{Name: "Ethiopia"},
	{Name: "Philippines"},
	{Name: "Egypt"},
	{Name: "Vietnam"},
	{Name: "Congo"},
	{Name: "Turkey"},
	{Name: "Iran"},
	{Name: "Germany"},
	{Name: "France"},
	{Name: "United Kingdom"},
	{Name: "Thailand"},
	{Name: "Italy"},
	{Name: "Tanzania"},
	{Name: "South Africa"},
	{Name: "Myanmar"},
	{Name: "South Korea"},
	{Name: "Kenya"},
	{Name: "Colombia"},
	{Name: "Spain"},
	{Name: "Argentina"},
	{Name: "Algeria"},
	{Name: "Sudan"},
	{Name: "Ukraine"},
	{Name: "Uganda"},
	{Name: "Canada"},
	{Name: "Iraq"},
	{Name: "Morocco"},
	{Name: "Afghanistan"},
	{Name: "Poland"},
	{Name: "Malaysia"},
	{Name: "Venezuela"},
	{Name: "Peru"},
	{Name: "Uzbekistan"},
	{Name: "Saudi Arabia"},
	{Name: "Mozambique"},
	{Name: "Ghana"},
	{Name: "Yemen"},
	{Name: "Nepal"},
	{Name: "Venezuela"},
	{Name: "Madagascar"},
	{Name: "Cameroon"},
	{Name: "Ivory Coast"},
	{Name: "Niger"},
	{Name: "Mali"},
	{Name: "Malawi"},
	{Name: "Kazakhstan"},
	{Name: "Ecuador"},
	{Name: "Guatemala"},
	{Name: "Netherlands"},
	{Name: "Syria"},
	{Name: "Cambodia"},
	{Name: "Senegal"},
	{Name: "Chad"},
	{Name: "Somalia"},
	{Name: "Zimbabwe"},
	{Name: "Rwanda"},
	{Name: "Benin"},
	{Name: "Burundi"},
	{Name: "Tunisia"},
	{Name: "Bolivia"},
	{Name: "Belgium"},
	{Name: "Czech Republic"},
	{Name: "Greece"},
	{Name: "Portugal"},
	{Name: "Sweden"},
	{Name: "Azerbaijan"},
	{Name: "Hungary"},
	{Name: "Belarus"},
	{Name: "Austria"},
	{Name: "Switzerland"},
	{Name: "Bolivia"},
	{Name: "Jordan"},
	{Name: "Serbia"},
	{Name: "Libya"},
	{Name: "United Arab Emirates"},
	{Name: "Honduras"},
	{Name: "Lebanon"},
	{Name: "Burma"},
	{Name: "Lebanon"},
	{Name: "Singapore"},
	{Name: "Belarus"},
	{Name: "Sri Lanka"},
	{Name: "Luxembourg"},
	{Name: "Kyrgyzstan"},
	{Name: "Turkmenistan"},
	{Name: "Sierra Leone"},
	{Name: "Tajikistan"},
	{Name: "Cyprus"},
	{Name: "Palestine"},
}

func (c *Country) save() error {
	if err := DB.Create(&c).Error; err != nil {
		return fmt.Errorf("DB.Create: %w", err)
	}
	return nil
}

func AddCountries() error {
	for _, country := range CountryList {
		if err := country.save(); err != nil {
			return fmt.Errorf("country.Save: %w", err)
		}
	}
	return nil
}

func GetAllCountries() (*[]Country, error) {
	var c []Country

	if err := DB.Model(&[]Country{}).Scan(&c).Error; err != nil {
		return &c, fmt.Errorf("DB.Model(&[]Country{}).Scan: %w", err)
	}
	return &c, nil
}
