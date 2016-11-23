package main

import (
	"io/ioutil"
	"os"
	"testing"

	. "github.com/onsi/gomega"
	"gopkg.in/yaml.v2"
)

var (
	defaultHoverflyHost      = "localhost"
	defaultHoverflyAdminPort = "8888"
	defaultHoverflyProxyPort = "8500"
	defaultHoverflyUsername  = ""
	defaultHoverflyPassword  = ""
)

func Test_GetConfigWillReturnTheDefaultValues(t *testing.T) {
	RegisterTestingT(t)

	SetConfigurationDefaults()
	result := GetConfig()

	Expect(result.HoverflyHost).To(Equal(defaultHoverflyHost))
	Expect(result.HoverflyAdminPort).To(Equal(defaultHoverflyAdminPort))
	Expect(result.HoverflyProxyPort).To(Equal(defaultHoverflyProxyPort))
	Expect(result.HoverflyUsername).To(Equal(defaultHoverflyUsername))
	Expect(result.HoverflyPassword).To(Equal(defaultHoverflyPassword))
}

func Test_Config_SetHost_OverridesDefaultValueWithAHoverflyHost(t *testing.T) {
	RegisterTestingT(t)

	SetConfigurationDefaults()
	result := GetConfig().SetHost("testhost")

	Expect(result.HoverflyHost).To(Equal("testhost"))
	Expect(result.HoverflyAdminPort).To(Equal(defaultHoverflyAdminPort))
	Expect(result.HoverflyProxyPort).To(Equal(defaultHoverflyProxyPort))
	Expect(result.HoverflyUsername).To(Equal(defaultHoverflyUsername))
	Expect(result.HoverflyPassword).To(Equal(defaultHoverflyPassword))
}

func Test_Config_SetHost_DoesNotOverrideWhenEmpty(t *testing.T) {
	RegisterTestingT(t)

	SetConfigurationDefaults()
	result := GetConfig().SetHost("")

	Expect(result.HoverflyHost).To(Equal(defaultHoverflyHost))
	Expect(result.HoverflyAdminPort).To(Equal(defaultHoverflyAdminPort))
	Expect(result.HoverflyProxyPort).To(Equal(defaultHoverflyProxyPort))
	Expect(result.HoverflyUsername).To(Equal(defaultHoverflyUsername))
	Expect(result.HoverflyPassword).To(Equal(defaultHoverflyPassword))
}

func Test_Config_SetAdminPort_OverridesDefaultValueWithAHoverflyAdminPort(t *testing.T) {
	RegisterTestingT(t)

	SetConfigurationDefaults()
	result := GetConfig().SetAdminPort("5")

	Expect(result.HoverflyHost).To(Equal(defaultHoverflyHost))
	Expect(result.HoverflyAdminPort).To(Equal("5"))
	Expect(result.HoverflyProxyPort).To(Equal(defaultHoverflyProxyPort))
	Expect(result.HoverflyUsername).To(Equal(defaultHoverflyUsername))
	Expect(result.HoverflyPassword).To(Equal(defaultHoverflyPassword))
}

func Test_Config_SetAdminPort_DoesNotOverrideWhenEmpty(t *testing.T) {
	RegisterTestingT(t)

	SetConfigurationDefaults()
	result := GetConfig().SetAdminPort("")

	Expect(result.HoverflyHost).To(Equal(defaultHoverflyHost))
	Expect(result.HoverflyAdminPort).To(Equal(defaultHoverflyAdminPort))
	Expect(result.HoverflyProxyPort).To(Equal(defaultHoverflyProxyPort))
	Expect(result.HoverflyUsername).To(Equal(defaultHoverflyUsername))
	Expect(result.HoverflyPassword).To(Equal(defaultHoverflyPassword))
}

func Test_Config_SetProxyPort_OverridesDefaultValueWithAHoverflyProxyPort(t *testing.T) {
	RegisterTestingT(t)

	SetConfigurationDefaults()
	result := GetConfig().SetProxyPort("7")

	Expect(result.HoverflyHost).To(Equal(defaultHoverflyHost))
	Expect(result.HoverflyAdminPort).To(Equal(defaultHoverflyAdminPort))
	Expect(result.HoverflyProxyPort).To(Equal("7"))
	Expect(result.HoverflyUsername).To(Equal(defaultHoverflyUsername))
	Expect(result.HoverflyPassword).To(Equal(defaultHoverflyPassword))
}

func Test_Config_SetProxyPort_DoesNotOverrideWhenEmpty(t *testing.T) {
	RegisterTestingT(t)

	SetConfigurationDefaults()
	result := GetConfig().SetProxyPort("")

	Expect(result.HoverflyHost).To(Equal(defaultHoverflyHost))
	Expect(result.HoverflyAdminPort).To(Equal(defaultHoverflyAdminPort))
	Expect(result.HoverflyProxyPort).To(Equal(defaultHoverflyProxyPort))
	Expect(result.HoverflyUsername).To(Equal(defaultHoverflyUsername))
	Expect(result.HoverflyPassword).To(Equal(defaultHoverflyPassword))
}

func Test_Config_SetUsername_OverridesDefaultValueWithAHoverflyUsername(t *testing.T) {
	RegisterTestingT(t)

	SetConfigurationDefaults()
	result := GetConfig().SetUsername("benjih")

	Expect(result.HoverflyHost).To(Equal(defaultHoverflyHost))
	Expect(result.HoverflyAdminPort).To(Equal(defaultHoverflyAdminPort))
	Expect(result.HoverflyProxyPort).To(Equal(defaultHoverflyProxyPort))
	Expect(result.HoverflyUsername).To(Equal("benjih"))
	Expect(result.HoverflyPassword).To(Equal(defaultHoverflyPassword))
}

func Test_Config_SetUsername_DoesNotOverrideWhenEmpty(t *testing.T) {
	RegisterTestingT(t)

	SetConfigurationDefaults()
	result := GetConfig().SetUsername("")

	Expect(result.HoverflyHost).To(Equal(defaultHoverflyHost))
	Expect(result.HoverflyAdminPort).To(Equal(defaultHoverflyAdminPort))
	Expect(result.HoverflyProxyPort).To(Equal(defaultHoverflyProxyPort))
	Expect(result.HoverflyUsername).To(Equal(defaultHoverflyUsername))
	Expect(result.HoverflyPassword).To(Equal(defaultHoverflyPassword))
}

func Test_Config_SetPassword_OverridesDefaultValueWithAHoverflyPassword(t *testing.T) {
	RegisterTestingT(t)

	SetConfigurationDefaults()
	result := GetConfig().SetPassword("burger-toucher")

	Expect(result.HoverflyHost).To(Equal(defaultHoverflyHost))
	Expect(result.HoverflyAdminPort).To(Equal(defaultHoverflyAdminPort))
	Expect(result.HoverflyProxyPort).To(Equal(defaultHoverflyProxyPort))
	Expect(result.HoverflyUsername).To(Equal(defaultHoverflyUsername))
	Expect(result.HoverflyPassword).To(Equal("burger-toucher"))
}

func Test_Config_SetPassword_DoesNotOverrideWhenEmpty(t *testing.T) {
	RegisterTestingT(t)

	SetConfigurationDefaults()
	result := GetConfig().SetPassword("")

	Expect(result.HoverflyHost).To(Equal(defaultHoverflyHost))
	Expect(result.HoverflyAdminPort).To(Equal(defaultHoverflyAdminPort))
	Expect(result.HoverflyProxyPort).To(Equal(defaultHoverflyProxyPort))
	Expect(result.HoverflyUsername).To(Equal(defaultHoverflyUsername))
	Expect(result.HoverflyPassword).To(Equal(defaultHoverflyPassword))
}

func Test_Config_WriteToFile_WritesTheConfigObjectToAFileInAYamlFormat(t *testing.T) {
	RegisterTestingT(t)

	SetConfigurationDefaults()
	config := GetConfig()
	config = config.SetHost("testhost").SetAdminPort("1234").SetProxyPort("4567").SetUsername("username").SetPassword("password")

	wd, _ := os.Getwd()
	hoverflyDirectory := HoverflyDirectory{
		Path: wd,
	}

	err := config.WriteToFile(hoverflyDirectory)

	Expect(err).To(BeNil())

	data, _ := ioutil.ReadFile(hoverflyDirectory.Path + "/config.yaml")
	os.Remove(hoverflyDirectory.Path + "/config.yaml")

	var result Config
	yaml.Unmarshal(data, &result)

	Expect(result.HoverflyHost).To(Equal("testhost"))
	Expect(result.HoverflyAdminPort).To(Equal("1234"))
	Expect(result.HoverflyProxyPort).To(Equal("4567"))
	Expect(result.HoverflyUsername).To(Equal("username"))
	Expect(result.HoverflyPassword).To(Equal("password"))
}
