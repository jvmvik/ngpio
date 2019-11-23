package ngpio

import (
	"errors"
	"log"
	"os"
	"sync"
	"io/ioutil"
"fmt"
)

var (
	lock sync.Mutex
	path string
)

// Port definiton
type Port struct {
	LogicalID           int
	PhysicalOutput      int
	PhysicalInput       int
	PhysicalOutputLabel string
	PhysicalInputLabel  string
}

// PortCollection store all the port definition per device
type PortSpecification struct {
	CardName    string
	CardVersion string
	ports       []Port
}

func (port Port) Output() {
	port.init("out")
}

func (port Port) init(direction string) {
	lock.Lock()
	defer lock.Unlock()
	if !present(path) {
		bb := make([]byte, 1)
		bb[0] = byte(port.LogicalID)
		err := ioutil.WriteFile("/sys/class/gpio/export", bb, 0666)
		if err != nil {
			log.Fatal("fail to add register port")
			os.Exit(1)
		}
	}
	err := ioutil.WriteFile(port.path()+"/direction", []byte(direction), 0666)
	if err != nil {
		log.Fatal("fail to set the direction")
		os.Exit(1)
	}
}

func (port Port) path() string {
	return fmt.Sprintf("/sys/class/gpio/gpio%d", port.LogicalID)
}

func (port Port) Input() {
	log.Fatalln("input port are not supported")
}

func (port Port) High() {
	port.value("1")
}

func (port Port) Low() {
	port.value("0")
}

func (port Port) value(state string) {
	v := []byte(state)
	err := ioutil.WriteFile(port.path()+"/value", v, 0666)
	if err != nil {
		log.Fatal("fail to set the value")
		os.Exit(1)
	}
}

// FindPortByPhysical port
func (spec *PortSpecification) FindPortByOutput(physicalOutput int) (Port, error) {
	for _, port := range spec.ports {
		if port.PhysicalOutput == physicalOutput {
			return port, nil
		}
	}
	return Port{}, errors.New("No port found")
}

// FindPort
func (spec *PortSpecification) FindPort(logicalID int) (Port, error) {
	for _, port := range spec.ports {
		if port.LogicalID == logicalID {
			return port, nil
		}
	}
	return Port{}, errors.New("No port found")
}

func present(filename string) bool {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

// Nano return PinDefCollection for Jetson Nano
func nano() PortSpecification {
	spec := PortSpecification{}
	spec.ports = []Port{
		Port{
			216,
			7,
			4,
			"GPIO9",
			"AUD_MCLK",
		},
		Port{
			50,
			11,
			17,
			"UART1_RTS",
			"UART2_RTS",
		},
		Port{
			79,
			12,
			18,
			"I2S0_SCLK",
			"DAP4_SCLK",
		},
		Port{
			14,
			13,
			27,
			"SPI1_SCK",
			"SPI2_SCK",
		},
		Port{
			194,
			15,
			22,
			"GPIO12",
			"LCD_TE",
		},
		Port{
			232,
			16,
			23,
			"SPI1_CS1",
			"SPI2_CS1",
		},
		Port{
			15,
			18,
			24,
			"SPI1_CS0",
			"SPI2_CS0",
		},
		Port{
			16,
			19,
			10,
			"SPI0_MOSI",
			"SPI1_MOSI",
		},
		Port{
			17,
			21,
			9,
			"SPI0_MISO",
			"SPI1_MISO",
		},
		Port{
			13,
			22,
			25,
			"SPI1_MISO",
			"SPI2_MISO",
		},
		Port{
			18,
			23,
			11,
			"SPI0_SCK",
			"SPI1_SCK",
		},
		Port{
			19,
			24,
			8,
			"SPI0_CS0",
			"SPI1_CS0",
		},
		Port{
			20,
			26,
			7,
			"SPI0_CS1",
			"SPI1_CS1",
		},
		Port{
			149,
			29,
			5,
			"GPIO01",
			"CAM_AF_EN",
		},
		Port{
			200,
			31,
			6,
			"GPIO11",
			"GPIO_PZ0",
		},
		Port{
			168,
			32,
			12,
			"GPIO07",
			"LCD_BL_PW",
			// "/sys/devices/7000a000.pwm",
			// 0
		},
		Port{
			38,
			33,
			13,
			"GPIO13",
			"GPIO_PE6",
			// "/sys/devices/7000a000.pwm",
			// 2
		},
		Port{
			76,
			35,
			19,
			"I2S0_FS",
			"DAP4_FS",
		},
		Port{
			51,
			36,
			16,
			"UART1_CTS",
			"UART2_CTS",
		},
		Port{
			12,
			37,
			26,
			"SPI1_MOSI",
			"SPI2_MOSI",
		},
		Port{
			77,
			38,
			20,
			"I2S0_DIN",
			"DAP4_DIN",
		},
		Port{
			78,
			40,
			21,
			"I2S0_DOUT",
			"DAP4_DOUT",
		},
	}
	return spec
}
