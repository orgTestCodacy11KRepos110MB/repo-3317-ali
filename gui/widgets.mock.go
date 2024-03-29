// Code generated by MockGen. DO NOT EDIT.
// Source: ./gui/widgets.go

// Package gui is a generated GoMock package.
package gui

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	canvas "github.com/mum4k/termdash/private/canvas"
	terminalapi "github.com/mum4k/termdash/terminal/terminalapi"
	widgetapi "github.com/mum4k/termdash/widgetapi"
	gauge "github.com/mum4k/termdash/widgets/gauge"
	linechart "github.com/mum4k/termdash/widgets/linechart"
	text "github.com/mum4k/termdash/widgets/text"
)

// MockLineChart is a mock of LineChart interface.
type MockLineChart struct {
	ctrl     *gomock.Controller
	recorder *MockLineChartMockRecorder
}

// MockLineChartMockRecorder is the mock recorder for MockLineChart.
type MockLineChartMockRecorder struct {
	mock *MockLineChart
}

// NewMockLineChart creates a new mock instance.
func NewMockLineChart(ctrl *gomock.Controller) *MockLineChart {
	mock := &MockLineChart{ctrl: ctrl}
	mock.recorder = &MockLineChartMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLineChart) EXPECT() *MockLineChartMockRecorder {
	return m.recorder
}

// Draw mocks base method.
func (m *MockLineChart) Draw(cvs *canvas.Canvas, meta *widgetapi.Meta) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Draw", cvs, meta)
	ret0, _ := ret[0].(error)
	return ret0
}

// Draw indicates an expected call of Draw.
func (mr *MockLineChartMockRecorder) Draw(cvs, meta interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Draw", reflect.TypeOf((*MockLineChart)(nil).Draw), cvs, meta)
}

// Keyboard mocks base method.
func (m *MockLineChart) Keyboard(k *terminalapi.Keyboard, meta *widgetapi.EventMeta) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Keyboard", k, meta)
	ret0, _ := ret[0].(error)
	return ret0
}

// Keyboard indicates an expected call of Keyboard.
func (mr *MockLineChartMockRecorder) Keyboard(k, meta interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Keyboard", reflect.TypeOf((*MockLineChart)(nil).Keyboard), k, meta)
}

// Mouse mocks base method.
func (m_2 *MockLineChart) Mouse(m *terminalapi.Mouse, meta *widgetapi.EventMeta) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Mouse", m, meta)
	ret0, _ := ret[0].(error)
	return ret0
}

// Mouse indicates an expected call of Mouse.
func (mr *MockLineChartMockRecorder) Mouse(m, meta interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Mouse", reflect.TypeOf((*MockLineChart)(nil).Mouse), m, meta)
}

// Options mocks base method.
func (m *MockLineChart) Options() widgetapi.Options {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Options")
	ret0, _ := ret[0].(widgetapi.Options)
	return ret0
}

// Options indicates an expected call of Options.
func (mr *MockLineChartMockRecorder) Options() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Options", reflect.TypeOf((*MockLineChart)(nil).Options))
}

// Series mocks base method.
func (m *MockLineChart) Series(label string, values []float64, opts ...linechart.SeriesOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{label, values}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Series", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Series indicates an expected call of Series.
func (mr *MockLineChartMockRecorder) Series(label, values interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{label, values}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Series", reflect.TypeOf((*MockLineChart)(nil).Series), varargs...)
}

// MockText is a mock of Text interface.
type MockText struct {
	ctrl     *gomock.Controller
	recorder *MockTextMockRecorder
}

// MockTextMockRecorder is the mock recorder for MockText.
type MockTextMockRecorder struct {
	mock *MockText
}

// NewMockText creates a new mock instance.
func NewMockText(ctrl *gomock.Controller) *MockText {
	mock := &MockText{ctrl: ctrl}
	mock.recorder = &MockTextMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockText) EXPECT() *MockTextMockRecorder {
	return m.recorder
}

// Draw mocks base method.
func (m *MockText) Draw(cvs *canvas.Canvas, meta *widgetapi.Meta) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Draw", cvs, meta)
	ret0, _ := ret[0].(error)
	return ret0
}

// Draw indicates an expected call of Draw.
func (mr *MockTextMockRecorder) Draw(cvs, meta interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Draw", reflect.TypeOf((*MockText)(nil).Draw), cvs, meta)
}

// Keyboard mocks base method.
func (m *MockText) Keyboard(k *terminalapi.Keyboard, meta *widgetapi.EventMeta) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Keyboard", k, meta)
	ret0, _ := ret[0].(error)
	return ret0
}

// Keyboard indicates an expected call of Keyboard.
func (mr *MockTextMockRecorder) Keyboard(k, meta interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Keyboard", reflect.TypeOf((*MockText)(nil).Keyboard), k, meta)
}

// Mouse mocks base method.
func (m_2 *MockText) Mouse(m *terminalapi.Mouse, meta *widgetapi.EventMeta) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Mouse", m, meta)
	ret0, _ := ret[0].(error)
	return ret0
}

// Mouse indicates an expected call of Mouse.
func (mr *MockTextMockRecorder) Mouse(m, meta interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Mouse", reflect.TypeOf((*MockText)(nil).Mouse), m, meta)
}

// Options mocks base method.
func (m *MockText) Options() widgetapi.Options {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Options")
	ret0, _ := ret[0].(widgetapi.Options)
	return ret0
}

// Options indicates an expected call of Options.
func (mr *MockTextMockRecorder) Options() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Options", reflect.TypeOf((*MockText)(nil).Options))
}

// Write mocks base method.
func (m *MockText) Write(text string, wOpts ...text.WriteOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{text}
	for _, a := range wOpts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Write", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Write indicates an expected call of Write.
func (mr *MockTextMockRecorder) Write(text interface{}, wOpts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{text}, wOpts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockText)(nil).Write), varargs...)
}

// MockGauge is a mock of Gauge interface.
type MockGauge struct {
	ctrl     *gomock.Controller
	recorder *MockGaugeMockRecorder
}

// MockGaugeMockRecorder is the mock recorder for MockGauge.
type MockGaugeMockRecorder struct {
	mock *MockGauge
}

// NewMockGauge creates a new mock instance.
func NewMockGauge(ctrl *gomock.Controller) *MockGauge {
	mock := &MockGauge{ctrl: ctrl}
	mock.recorder = &MockGaugeMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGauge) EXPECT() *MockGaugeMockRecorder {
	return m.recorder
}

// Draw mocks base method.
func (m *MockGauge) Draw(cvs *canvas.Canvas, meta *widgetapi.Meta) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Draw", cvs, meta)
	ret0, _ := ret[0].(error)
	return ret0
}

// Draw indicates an expected call of Draw.
func (mr *MockGaugeMockRecorder) Draw(cvs, meta interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Draw", reflect.TypeOf((*MockGauge)(nil).Draw), cvs, meta)
}

// Keyboard mocks base method.
func (m *MockGauge) Keyboard(k *terminalapi.Keyboard, meta *widgetapi.EventMeta) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Keyboard", k, meta)
	ret0, _ := ret[0].(error)
	return ret0
}

// Keyboard indicates an expected call of Keyboard.
func (mr *MockGaugeMockRecorder) Keyboard(k, meta interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Keyboard", reflect.TypeOf((*MockGauge)(nil).Keyboard), k, meta)
}

// Mouse mocks base method.
func (m_2 *MockGauge) Mouse(m *terminalapi.Mouse, meta *widgetapi.EventMeta) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Mouse", m, meta)
	ret0, _ := ret[0].(error)
	return ret0
}

// Mouse indicates an expected call of Mouse.
func (mr *MockGaugeMockRecorder) Mouse(m, meta interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Mouse", reflect.TypeOf((*MockGauge)(nil).Mouse), m, meta)
}

// Options mocks base method.
func (m *MockGauge) Options() widgetapi.Options {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Options")
	ret0, _ := ret[0].(widgetapi.Options)
	return ret0
}

// Options indicates an expected call of Options.
func (mr *MockGaugeMockRecorder) Options() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Options", reflect.TypeOf((*MockGauge)(nil).Options))
}

// Percent mocks base method.
func (m *MockGauge) Percent(p int, opts ...gauge.Option) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{p}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Percent", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Percent indicates an expected call of Percent.
func (mr *MockGaugeMockRecorder) Percent(p interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{p}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Percent", reflect.TypeOf((*MockGauge)(nil).Percent), varargs...)
}
