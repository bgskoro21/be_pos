package generator

import (
	"fmt"
	"os"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func GenerateRepository(name string){
	capitalized := cases.Title(language.English).String(name)
	basePath := fmt.Sprintf("../repository/%s", name);
	os.MkdirAll(basePath, 0755)

	interfaceContent := fmt.Sprintf(`package repository

type %sRepository interface {
	// define methods
}
`, capitalized)

	implContent := fmt.Sprintf(`package repository

import "gorm.io/gorm"

type %sRepositoryImpl struct {
	db *gorm.DB
}

func New%sRepositoryImpl(db *gorm.DB) %sRepository {
	return &%sRepositoryImpl{db}
}
`, capitalized, capitalized, capitalized, capitalized)

	write(fmt.Sprintf("%s/%s_repository.go", basePath, name), interfaceContent)
	write(fmt.Sprintf("%s/%s_repository_impl.go", basePath, name), implContent)

	fmt.Printf("✅ Repository created for: %s\n", name)
}