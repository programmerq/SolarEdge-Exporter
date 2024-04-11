/*
MIT License

# Copyright (c) 2019 David Suarez

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/
package exporter

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	SunSpec_DID = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "SunSpec_DID",
		Help: "101 = single phase 102 = split phase1 103 = three phase",
	}, []string{"serial", "model", "version"})

	SunSpec_Length = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "SunSpec_Length",
		Help: "Registers 50 = Length of model block",
	}, []string{"serial", "model", "version"})

	AC_Current = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "AC_Current",
		Help: "Amps AC Total Current value",
	}, []string{"serial", "model", "version"})

	AC_CurrentA = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "AC_CurrentA",
		Help: "Amps AC Phase A Current value",
	}, []string{"serial", "model", "version"})

	AC_CurrentB = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "AC_CurrentB",
		Help: "Amps AC Phase B Current value",
	}, []string{"serial", "model", "version"})

	AC_CurrentC = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "AC_CurrentC",
		Help: "Amps AC Phase C Current value",
	}, []string{"serial", "model", "version"})

	AC_Current_SF = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "AC_Current_SF",
		Help: "AC Current scale factor",
	}, []string{"serial", "model", "version"})

	AC_VoltageAB = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "AC_VoltageAB",
		Help: "Volts AC Voltage Phase AB value",
	}, []string{"serial", "model", "version"})

	AC_VoltageBC = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "AC_VoltageBC",
		Help: "Volts AC Voltage Phase BC value",
	}, []string{"serial", "model", "version"})

	AC_VoltageCA = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "AC_VoltageCA",
		Help: "Volts AC Voltage Phase CA value",
	}, []string{"serial", "model", "version"})

	AC_VoltageAN = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "AC_VoltageAN",
		Help: "Volts AC Voltage Phase A to N value",
	}, []string{"serial", "model", "version"})

	AC_VoltageBN = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "AC_VoltageBN",
		Help: "Volts AC Voltage Phase B to N value",
	}, []string{"serial", "model", "version"})

	AC_VoltageCN = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "AC_VoltageCN",
		Help: "Volts AC Voltage Phase C to N value",
	}, []string{"serial", "model", "version"})

	AC_Voltage_SF = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "AC_Voltage_SF",
		Help: "AC Voltage scale factor",
	}, []string{"serial", "model", "version"})

	AC_Power = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "AC_Power",
		Help: "Watts AC Power value",
	}, []string{"serial", "model", "version"})

	AC_Power_SF = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "AC_Power_SF",
		Help: "AC Power scale factor",
	}, []string{"serial", "model", "version"})

	AC_Frequency = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "AC_Frequency",
		Help: "Hertz AC Frequency value",
	}, []string{"serial", "model", "version"})

	AC_Frequency_SF = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "AC_Frequency_SF",
		Help: "Scale factor",
	}, []string{"serial", "model", "version"})

	AC_VA = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "AC_VA",
		Help: "VA Apparent Power",
	}, []string{"serial", "model", "version"})

	AC_VA_SF = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "AC_VA_SF",
		Help: "Scale factor",
	}, []string{"serial", "model", "version"})

	AC_VAR = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "AC_VAR",
		Help: "VAR Reactive Power",
	}, []string{"serial", "model", "version"})

	AC_VAR_SF = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "AC_VAR_SF",
		Help: "Scale factor",
	}, []string{"serial", "model", "version"})

	AC_PF = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "AC_PF",
		Help: "% Power Factor",
	}, []string{"serial", "model", "version"})

	AC_PF_SF = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "AC_PF_SF",
		Help: "Scale factor",
	}, []string{"serial", "model", "version"})

	AC_Energy_WH = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "AC_Energy_WH",
		Help: "WattHours AC Lifetime Energy production",
	}, []string{"serial", "model", "version"})

	AC_Energy_WH_SF = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "AC_Energy_WH_SF",
		Help: "Scale factor",
	}, []string{"serial", "model", "version"})

	DC_Current = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "DC_Current",
		Help: "Amps DC Current value",
	}, []string{"serial", "model", "version"})

	DC_Current_SF = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "DC_Current_SF",
		Help: "Scale factor",
	}, []string{"serial", "model", "version"})

	DC_Voltage = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "DC_Voltage",
		Help: "Volts DC Voltage value",
	}, []string{"serial", "model", "version"})

	DC_Voltage_SF = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "DC_Voltage_SF",
		Help: "Scale factor",
	}, []string{"serial", "model", "version"})

	DC_Power = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "DC_Power",
		Help: "Watts DC Power value",
	}, []string{"serial", "model", "version"})

	DC_Power_SF = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "DC_Power_SF",
		Help: "Scale factor",
	}, []string{"serial", "model", "version"})

	Temp_Sink = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "Temp_Sink",
		Help: "Degrees C Heat Sink Temperature",
	}, []string{"serial", "model", "version"})

	Temp_SF = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "Temp_SF",
		Help: "Scale factor",
	}, []string{"serial", "model", "version"})

	Status = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "Status",
		Help: "Operating State",
	}, []string{"serial", "model", "version"})

	Status_Vendor = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "Status_Vendor",
		Help: "Vendor-defined operating state and error codes. For error description, meaning and troubleshooting, refer to the SolarEdge Installation Guide.",
	}, []string{"serial", "model", "version"})

	// Meter 1

	M_SunSpec_DID = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "M_SunSpec_DID",
		Help: "",
	}, []string{"serial", "model", "version"})

	M_SunSpec_Length = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "M_SunSpec_Length",
		Help: "",
	}, []string{"serial", "model", "version"})

	M_AC_Current = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "M_AC_Current",
		Help: "Amps AC Total Current value",
	}, []string{"serial", "model", "version"})

	M_AC_CurrentA = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "M_AC_CurrentA",
		Help: "Amps AC Phase A Current value",
	}, []string{"serial", "model", "version"})

	M_AC_CurrentB = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "M_AC_CurrentB",
		Help: "Amps AC Phase B Current value",
	}, []string{"serial", "model", "version"})

	M_AC_CurrentC = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "M_AC_CurrentC",
		Help: "Amps AC Phase C Current value",
	}, []string{"serial", "model", "version"})

	M_AC_Current_SF = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "M_AC_Current_SF",
		Help: "AC Current scale factor",
	}, []string{"serial", "model", "version"})

	M_AC_VoltageLN = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "M_AC_VoltageLN",
		Help: "Volts AC Voltage Phase AB value",
	}, []string{"serial", "model", "version"})

	M_AC_VoltageAN = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "M_AC_VoltageAN",
		Help: "Volts AC Voltage Phase BC value",
	}, []string{"serial", "model", "version"})

	M_AC_VoltageBN = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "M_AC_VoltageBN",
		Help: "Volts AC Voltage Phase BC value",
	}, []string{"serial", "model", "version"})

	M_AC_VoltageCN = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "M_AC_VoltageCN",
		Help: "Volts AC Voltage Phase BC value",
	}, []string{"serial", "model", "version"})

	M_AC_VoltageLL = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "M_AC_VoltageLL",
		Help: "Volts AC Voltage Phase BC value",
	}, []string{"serial", "model", "version"})

	M_AC_VoltageAB = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "M_AC_VoltageAB",
		Help: "Volts AC Voltage Phase BC value",
	}, []string{"serial", "model", "version"})

	M_AC_VoltageBC = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "M_AC_VoltageBC",
		Help: "Volts AC Voltage Phase BC value",
	}, []string{"serial", "model", "version"})

	M_AC_VoltageCA = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "M_AC_VoltageCA",
		Help: "Volts AC Voltage Phase BC value",
	}, []string{"serial", "model", "version"})

	M_AC_Voltage_SF = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "M_AC_Voltage_SF",
		Help: "AC Voltage scale factor",
	}, []string{"serial", "model", "version"})

	M_AC_Frequency = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "M_AC_Frequency",
		Help: "Hertz AC Frequency value",
	}, []string{"serial", "model", "version"})

	M_AC_Frequency_SF = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "M_AC_Frequency_SF",
		Help: "Scale factor",
	}, []string{"serial", "model", "version"})

	M_AC_Power = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "M_AC_Power",
		Help: "Watts AC Power value",
	}, []string{"serial", "model", "version"})

	M_AC_Power_A = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "M_AC_Power_A",
		Help: "Watts AC Power value",
	}, []string{"serial", "model", "version"})

	M_AC_Power_B = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "M_AC_Power_B",
		Help: "Watts AC Power value",
	}, []string{"serial", "model", "version"})

	M_AC_Power_C = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "M_AC_Power_C",
		Help: "Watts AC Power value",
	}, []string{"serial", "model", "version"})

	M_AC_Power_SF = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "M_AC_Power_SF",
		Help: "AC Power scale factor",
	}, []string{"serial", "model", "version"})

	M_AC_VA = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "M_AC_VA",
		Help: "VA Apparent Power",
	}, []string{"serial", "model", "version"})

	M_AC_VA_A = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "M_AC_VA_A",
		Help: "VA Apparent Power",
	}, []string{"serial", "model", "version"})

	M_AC_VA_B = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "M_AC_VA_B",
		Help: "VA Apparent Power",
	}, []string{"serial", "model", "version"})

	M_AC_VA_C = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "M_AC_VA_C",
		Help: "VA Apparent Power",
	}, []string{"serial", "model", "version"})

	M_AC_VA_SF = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "M_AC_VA_SF",
		Help: "Scale factor",
	}, []string{"serial", "model", "version"})

	M_AC_VAR = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "M_AC_VAR",
		Help: "VAR Reactive Power",
	}, []string{"serial", "model", "version"})

	M_AC_VAR_A = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "M_AC_VAR_A",
		Help: "VAR Reactive Power",
	}, []string{"serial", "model", "version"})

	M_AC_VAR_B = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "M_AC_VAR_B",
		Help: "VAR Reactive Power",
	}, []string{"serial", "model", "version"})

	M_AC_VAR_C = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "M_AC_VAR_C",
		Help: "VAR Reactive Power",
	}, []string{"serial", "model", "version"})

	M_AC_VAR_SF = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "M_AC_VAR_SF",
		Help: "Scale factor",
	}, []string{"serial", "model", "version"})

	M_AC_PF = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "M_AC_PF",
		Help: "% Power Factor",
	}, []string{"serial", "model", "version"})

	M_AC_PF_A = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "M_AC_PF_A",
		Help: "% Power Factor",
	}, []string{"serial", "model", "version"})

	M_AC_PF_B = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "M_AC_PF_B",
		Help: "% Power Factor",
	}, []string{"serial", "model", "version"})

	M_AC_PF_C = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "M_AC_PF_C",
		Help: "% Power Factor",
	}, []string{"serial", "model", "version"})

	M_AC_PF_SF = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "M_AC_PF_SF",
		Help: "Scale factor",
	}, []string{"serial", "model", "version"})

	M_Exported = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "M_Exported",
		Help: "WattHours AC Exported",
	}, []string{"serial", "model", "version"})

	M_Exported_A = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "M_Exported_A",
		Help: "WattHours AC Exported",
	}, []string{"serial", "model", "version"})

	M_Exported_B = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "M_Exported_B",
		Help: "WattHours AC Exported",
	}, []string{"serial", "model", "version"})

	M_Exported_C = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "M_Exported_C",
		Help: "WattHours AC Exported",
	}, []string{"serial", "model", "version"})

	M_Imported = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "M_Imported",
		Help: "WattHours AC Imported",
	}, []string{"serial", "model", "version"})

	M_Imported_A = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "M_Imported_A",
		Help: "WattHours AC Imported",
	}, []string{"serial", "model", "version"})

	M_Imported_B = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "M_Imported_B",
		Help: "WattHours AC Imported",
	}, []string{"serial", "model", "version"})

	M_Imported_C = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "M_Imported_C",
		Help: "WattHours AC Imported",
	}, []string{"serial", "model", "version"})

	M_Energy_W_SF = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "M_Energy_W_SF",
		Help: "M_Energy_W_SF",
	}, []string{"serial", "model", "version"})

	//Meter 2

	M2_SunSpec_DID = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "M2_SunSpec_DID",
		Help: "",
	}, []string{"serial", "model", "version"})

	M2_SunSpec_Length = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "M2_SunSpec_Length",
		Help: "",
	}, []string{"serial", "model", "version"})

	M2_AC_Current = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "M2_AC_Current",
		Help: "Amps AC Total Current value",
	}, []string{"serial", "model", "version"})

	M2_AC_CurrentA = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "M2_AC_CurrentA",
		Help: "Amps AC Phase A Current value",
	}, []string{"serial", "model", "version"})

	M2_AC_CurrentB = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "M2_AC_CurrentB",
		Help: "Amps AC Phase B Current value",
	}, []string{"serial", "model", "version"})

	M2_AC_CurrentC = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "M2_AC_CurrentC",
		Help: "Amps AC Phase C Current value",
	}, []string{"serial", "model", "version"})

	M2_AC_Current_SF = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "M2_AC_Current_SF",
		Help: "AC Current scale factor",
	}, []string{"serial", "model", "version"})

	M2_AC_VoltageLN = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "M2_AC_VoltageLN",
		Help: "Volts AC Voltage Phase AB value",
	}, []string{"serial", "model", "version"})

	M2_AC_VoltageAN = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "M2_AC_VoltageAN",
		Help: "Volts AC Voltage Phase BC value",
	}, []string{"serial", "model", "version"})

	M2_AC_VoltageBN = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "M2_AC_VoltageBN",
		Help: "Volts AC Voltage Phase BC value",
	}, []string{"serial", "model", "version"})

	M2_AC_VoltageCN = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "M2_AC_VoltageCN",
		Help: "Volts AC Voltage Phase BC value",
	}, []string{"serial", "model", "version"})

	M2_AC_VoltageLL = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "M2_AC_VoltageLL",
		Help: "Volts AC Voltage Phase BC value",
	}, []string{"serial", "model", "version"})

	M2_AC_VoltageAB = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "M2_AC_VoltageAB",
		Help: "Volts AC Voltage Phase BC value",
	}, []string{"serial", "model", "version"})

	M2_AC_VoltageBC = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "M2_AC_VoltageBC",
		Help: "Volts AC Voltage Phase BC value",
	}, []string{"serial", "model", "version"})

	M2_AC_VoltageCA = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "M2_AC_VoltageCA",
		Help: "Volts AC Voltage Phase BC value",
	}, []string{"serial", "model", "version"})

	M2_AC_Voltage_SF = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "M2_AC_Voltage_SF",
		Help: "AC Voltage scale factor",
	}, []string{"serial", "model", "version"})

	M2_AC_Frequency = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "M2_AC_Frequency",
		Help: "Hertz AC Frequency value",
	}, []string{"serial", "model", "version"})

	M2_AC_Frequency_SF = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "M2_AC_Frequency_SF",
		Help: "Scale factor",
	}, []string{"serial", "model", "version"})

	M2_AC_Power = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "M2_AC_Power",
		Help: "Watts AC Power value",
	}, []string{"serial", "model", "version"})

	M2_AC_Power_A = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "M2_AC_Power_A",
		Help: "Watts AC Power value",
	}, []string{"serial", "model", "version"})

	M2_AC_Power_B = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "M2_AC_Power_B",
		Help: "Watts AC Power value",
	}, []string{"serial", "model", "version"})

	M2_AC_Power_C = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "M2_AC_Power_C",
		Help: "Watts AC Power value",
	}, []string{"serial", "model", "version"})

	M2_AC_Power_SF = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "M2_AC_Power_SF",
		Help: "AC Power scale factor",
	}, []string{"serial", "model", "version"})

	M2_AC_VA = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "M2_AC_VA",
		Help: "VA Apparent Power",
	}, []string{"serial", "model", "version"})

	M2_AC_VA_A = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "M2_AC_VA_A",
		Help: "VA Apparent Power",
	}, []string{"serial", "model", "version"})

	M2_AC_VA_B = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "M2_AC_VA_B",
		Help: "VA Apparent Power",
	}, []string{"serial", "model", "version"})

	M2_AC_VA_C = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "M2_AC_VA_C",
		Help: "VA Apparent Power",
	}, []string{"serial", "model", "version"})

	M2_AC_VA_SF = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "M2_AC_VA_SF",
		Help: "Scale factor",
	}, []string{"serial", "model", "version"})

	M2_AC_VAR = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "M2_AC_VAR",
		Help: "VAR Reactive Power",
	}, []string{"serial", "model", "version"})

	M2_AC_VAR_A = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "M2_AC_VAR_A",
		Help: "VAR Reactive Power",
	}, []string{"serial", "model", "version"})

	M2_AC_VAR_B = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "M2_AC_VAR_B",
		Help: "VAR Reactive Power",
	}, []string{"serial", "model", "version"})

	M2_AC_VAR_C = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "M2_AC_VAR_C",
		Help: "VAR Reactive Power",
	}, []string{"serial", "model", "version"})

	M2_AC_VAR_SF = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "M2_AC_VAR_SF",
		Help: "Scale factor",
	}, []string{"serial", "model", "version"})

	M2_AC_PF = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "M2_AC_PF",
		Help: "% Power Factor",
	}, []string{"serial", "model", "version"})

	M2_AC_PF_A = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "M2_AC_PF_A",
		Help: "% Power Factor",
	}, []string{"serial", "model", "version"})

	M2_AC_PF_B = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "M2_AC_PF_B",
		Help: "% Power Factor",
	}, []string{"serial", "model", "version"})

	M2_AC_PF_C = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "M2_AC_PF_C",
		Help: "% Power Factor",
	}, []string{"serial", "model", "version"})

	M2_AC_PF_SF = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "M2_AC_PF_SF",
		Help: "Scale factor",
	}, []string{"serial", "model", "version"})

	M2_Exported = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "M2_Exported",
		Help: "WattHours AC Exported",
	}, []string{"serial", "model", "version"})

	M2_Exported_A = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "M2_Exported_A",
		Help: "WattHours AC Exported",
	}, []string{"serial", "model", "version"})

	M2_Exported_B = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "M2_Exported_B",
		Help: "WattHours AC Exported",
	}, []string{"serial", "model", "version"})

	M2_Exported_C = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "M2_Exported_C",
		Help: "WattHours AC Exported",
	}, []string{"serial", "model", "version"})

	M2_Imported = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "M2_Imported",
		Help: "WattHours AC Imported",
	}, []string{"serial", "model", "version"})

	M2_Imported_A = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "M2_Imported_A",
		Help: "WattHours AC Imported",
	}, []string{"serial", "model", "version"})

	M2_Imported_B = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "M2_Imported_B",
		Help: "WattHours AC Imported",
	}, []string{"serial", "model", "version"})

	M2_Imported_C = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "M2_Imported_C",
		Help: "WattHours AC Imported",
	}, []string{"serial", "model", "version"})

	M2_Energy_W_SF = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "M2_Energy_W_SF",
		Help: "M2_Energy_W_SF",
	}, []string{"serial", "model", "version"})
)
