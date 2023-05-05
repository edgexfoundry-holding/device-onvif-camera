// -*- Mode: Go; indent-tabs-mode: t -*-
//
// Copyright (C) 2022-2023 Intel Corporation
// Copyright (c) 2023 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package driver

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/IOTechSystems/onvif"
	"github.com/IOTechSystems/onvif/xsd"
	xsdOnvif "github.com/IOTechSystems/onvif/xsd/onvif"
	"github.com/stretchr/testify/mock"

	"github.com/IOTechSystems/onvif/device"
	sdkMocks "github.com/edgexfoundry/device-sdk-go/v3/pkg/interfaces/mocks"
	sdkModel "github.com/edgexfoundry/device-sdk-go/v3/pkg/models"

	"github.com/edgexfoundry/go-mod-core-contracts/v3/clients/logger"
	"github.com/edgexfoundry/go-mod-core-contracts/v3/errors"
	"github.com/edgexfoundry/go-mod-core-contracts/v3/models"
	contract "github.com/edgexfoundry/go-mod-core-contracts/v3/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	testDeviceName = "test-device"
	getFunction    = "getFunction"
)

var (
	ptrTrue  = boolPointer(true)
	ptrFalse = boolPointer(false)
)

func boolPointer(val bool) *xsd.Boolean {
	b := xsd.Boolean(val)
	return &b
}

func createDriverWithMockService() (*Driver, *sdkMocks.DeviceServiceSDK) {
	mockService := &sdkMocks.DeviceServiceSDK{}
	driver := NewDriver()
	driver.lc = logger.MockLogger{}
	driver.sdkService = mockService
	mockService.On("LoggingClient").Return(driver.lc).Maybe()
	return driver, mockService
}

func createTestDevice() models.Device {
	return models.Device{Name: testDeviceName, Protocols: map[string]models.ProtocolProperties{
		OnvifProtocol: map[string]interface{}{
			DeviceStatus: Unreachable,
		},
	}}
}

func createTestDeviceWithProtocols(protocols map[string]models.ProtocolProperties) models.Device {
	return models.Device{Name: testDeviceName, Protocols: protocols}
}

func TestDriver_HandleReadCommands(t *testing.T) {
	driver, mockService := createDriverWithMockService()

	tests := []struct {
		name          string
		deviceName    string
		protocols     map[string]models.ProtocolProperties
		reqs          []sdkModel.CommandRequest
		resp          string
		data          string
		expected      []*sdkModel.CommandValue
		errorExpected bool
	}{
		{
			name:       "simple read for RebootNeeded",
			deviceName: testDeviceName,
			reqs: []sdkModel.CommandRequest{
				{
					DeviceResourceName: RebootNeeded,
					Attributes: map[string]interface{}{
						getFunction: RebootNeeded,
						"service":   EdgeXWebService,
					},
					Type: "Bool",
				}},
			expected: []*sdkModel.CommandValue{
				{
					DeviceResourceName: RebootNeeded,
					Type:               "Bool",
					Value:              false,
					Tags:               map[string]string{},
				}},
		},
		{
			name:       "simple read of DeviceInformation",
			deviceName: testDeviceName,
			reqs: []sdkModel.CommandRequest{
				{
					DeviceResourceName: "DeviceInformation",
					Attributes: map[string]interface{}{
						getFunction: "GetDeviceInformation",
						"service":   onvif.DeviceWebService,
					},
					Type: "Object",
				}},
			resp: `<?xml version="1.0" encoding="UTF-8"?>
<Envelope xmlns="http://www.w3.org/2003/05/soap-envelope">
  <Header />
  <Body>
    <Content>
      <Manufacturer>Intel</Manufacturer>
      <Model>SimCamera</Model>
      <FirmwareVersion>2.4a</FirmwareVersion>
      <SerialNumber>46d1ab8d</SerialNumber>
      <HardwareId>1.0</HardwareId>
    </Content>
  </Body>
</Envelope>`,
			expected: []*sdkModel.CommandValue{
				{
					DeviceResourceName: "DeviceInformation",
					Type:               "Object",
					Value: &device.GetDeviceInformationResponse{
						Manufacturer:    "Intel",
						Model:           "SimCamera",
						FirmwareVersion: "2.4a",
						SerialNumber:    "46d1ab8d",
						HardwareId:      "1.0",
					},
					Tags: map[string]string{},
				}},
		},
		{
			name:       "simple read of GetNetworkInterfaces",
			deviceName: testDeviceName,
			reqs: []sdkModel.CommandRequest{
				{
					DeviceResourceName: "NetworkInterfaces",
					Attributes: map[string]interface{}{
						getFunction: "GetNetworkInterfaces",
						"service":   onvif.DeviceWebService,
					},
					Type: "Object",
				}},
			resp: `<?xml version="1.0" encoding="UTF-8"?>
<Envelope xmlns="http://www.w3.org/2003/05/soap-envelope">
  <Header />
  <Body>
    <Content>
      <NetworkInterfaces token="NET_TOKEN_4047201479">
        <Enabled>true</Enabled>
        <Info>
          <Name>eth0</Name>
          <HwAddress>02:42:C0:A8:90:0E</HwAddress>
          <MTU>1500</MTU>
        </Info>
        <IPv4>
          <Enabled>true</Enabled>
          <Config>
            <Manual>
              <Address>192.168.144.14</Address>
              <PrefixLength>20</PrefixLength>
            </Manual>
            <DHCP>false</DHCP>
          </Config>
        </IPv4>
      </NetworkInterfaces>
    </Content>
  </Body>
</Envelope>`,
			expected: []*sdkModel.CommandValue{
				{
					DeviceResourceName: "NetworkInterfaces",
					Type:               "Object",
					Value: &device.GetNetworkInterfacesResponse{
						NetworkInterfaces: xsdOnvif.NetworkInterface{
							DeviceEntity: xsdOnvif.DeviceEntity{
								Token: "NET_TOKEN_4047201479",
							},
							Enabled: ptrTrue,
							Info: &xsdOnvif.NetworkInterfaceInfo{
								Name:      "eth0",
								HwAddress: "02:42:C0:A8:90:0E",
								MTU:       1500,
							},
							IPv4: &xsdOnvif.IPv4NetworkInterface{
								Enabled: ptrTrue,
								Config: &xsdOnvif.IPv4Configuration{
									Manual: &xsdOnvif.PrefixedIPv4Address{
										Address:      "192.168.144.14",
										PrefixLength: 20,
									},
									DHCP: ptrFalse,
								},
							},
						},
					},
					Tags: map[string]string{},
				}},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
				_, err := writer.Write([]byte(test.resp))
				assert.NoError(t, err)
			}))
			defer server.Close()

			client, mockDevice := createOnvifClientWithMockDevice(driver, testDeviceName)
			driver.onvifClients = map[string]*OnvifClient{
				testDeviceName: client,
			}

			mockService.On("GetDeviceByName", testDeviceName).
				Return(createTestDevice(), nil)

			mockDevice.On("GetEndpointByRequestStruct", mock.Anything).Return(server.URL, nil)

			sendSoap := mockDevice.On("SendSoap", mock.Anything, mock.Anything)
			sendSoap.Run(func(args mock.Arguments) {
				resp, err := http.Post(server.URL, "application/soap+xml; charset=utf-8", strings.NewReader(args.String(1)))
				sendSoap.Return(resp, err)
			})

			actual, err := driver.HandleReadCommands(test.deviceName, test.protocols, test.reqs)
			if test.errorExpected {
				require.Error(t, err)
			}
			assert.Equal(t, test.expected, actual)
		})
	}
}

// TestUpdateDevice verifies proper updating of device information
func TestUpdateDevice(t *testing.T) {
	driver, mockService := createDriverWithMockService()
	tests := []struct {
		device  models.Device
		devInfo *device.GetDeviceInformationResponse

		expectedDevice           models.Device
		errorExpected            bool
		updateDeviceExpected     bool
		addDeviceExpected        bool
		removeDeviceExpected     bool
		removeDeviceFailExpected bool
	}{
		{
			device: contract.Device{
				Name: "testName",
			},
			updateDeviceExpected: true,
			devInfo: &device.GetDeviceInformationResponse{
				Manufacturer:    "Intel",
				Model:           "SimCamera",
				FirmwareVersion: "2.5a",
				SerialNumber:    "9a32410c",
				HardwareId:      "1.0",
			},
		},
		{
			removeDeviceExpected:     true,
			removeDeviceFailExpected: true,
			addDeviceExpected:        true,
			device: contract.Device{
				Name: "unknown_unknown_device",
				Protocols: map[string]models.ProtocolProperties{
					OnvifProtocol: map[string]interface{}{
						EndpointRefAddress: "793dfb2-28b0-11ed-a261-0242ac120002",
					},
				}},
			devInfo: &device.GetDeviceInformationResponse{
				Manufacturer:    "Intel",
				Model:           "SimCamera",
				FirmwareVersion: "2.5a",
				SerialNumber:    "9a32410c",
				HardwareId:      "1.0",
			},
			expectedDevice: contract.Device{
				Name: "Intel-SimCamera-793dfb2-28b0-11ed-a261-0242ac120002",
				Protocols: map[string]models.ProtocolProperties{
					OnvifProtocol: map[string]interface{}{
						EndpointRefAddress: "793dfb2-28b0-11ed-a261-0242ac120002",
					},
				},
			},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.device.Name, func(t *testing.T) {

			if test.removeDeviceExpected {
				if test.removeDeviceFailExpected {
					mockService.On("RemoveDeviceByName", test.device.Name).Return(errors.NewCommonEdgeX(errors.KindContractInvalid, "unit test error", nil)).Once()
				} else {
					mockService.On("RemoveDeviceByName", test.device.Name).Return(nil).Once()
				}
			}

			if test.updateDeviceExpected {
				mockService.On("UpdateDevice", test.device).Return(nil).Once()
			}

			if test.addDeviceExpected {
				mockService.On("AddDevice", test.expectedDevice).Return(test.expectedDevice.Name, nil).Once()
			}

			err := driver.updateDevice(test.device, test.devInfo)

			mockService.AssertExpectations(t)
			if test.errorExpected {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestDriver_RemoveDevice(t *testing.T) {
	driver, _ := createDriverWithMockService()
	driver.onvifClients = map[string]*OnvifClient{
		testDeviceName: {},
	}

	assert.Len(t, driver.onvifClients, 1)

	// remove a non-existent device. should be no error and still have 1 device
	err := driver.RemoveDevice("bogus device", map[string]models.ProtocolProperties{})
	require.NoError(t, err)
	assert.Len(t, driver.onvifClients, 1)

	// remove actual device and check that there are no devices left
	err = driver.RemoveDevice(testDeviceName, map[string]models.ProtocolProperties{})
	require.NoError(t, err)
	assert.Len(t, driver.onvifClients, 0)
}
