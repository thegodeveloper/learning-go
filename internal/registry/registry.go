package registry

import "sort"

// Module is the contract every learning package satisfies.
// Packages with no sub-functions return nil from Functions().
type Module interface {
	Name() string
	Run(show bool)
	Functions() map[string]func(bool)
}

var modules = map[string]Module{}

// Register is called from each package's init() in its register.go file.
func Register(m Module) {
	modules[m.Name()] = m
}

// Names returns all registered module names, sorted.
func Names() []string {
	names := make([]string, 0, len(modules))
	for name := range modules {
		names = append(names, name)
	}
	sort.Strings(names)
	return names
}

// LookupRun returns the Run function for a named module.
func LookupRun(name string) (func(bool), bool) {
	m, ok := modules[name]
	if !ok {
		return nil, false
	}
	return m.Run, true
}

// HasSubFunctions reports whether a module exposes named sub-functions.
func HasSubFunctions(name string) bool {
	m, ok := modules[name]
	return ok && len(m.Functions()) > 0
}

// GetFunctionNames returns sorted sub-function names for a module.
func GetFunctionNames(name string) []string {
	m, ok := modules[name]
	if !ok {
		return nil
	}
	fns := m.Functions()
	names := make([]string, 0, len(fns))
	for n := range fns {
		names = append(names, n)
	}
	sort.Strings(names)
	return names
}

// GetPackageFunctions returns the sub-function map for a module, or nil.
func GetPackageFunctions(name string) map[string]func(bool) {
	m, ok := modules[name]
	if !ok {
		return nil
	}
	return m.Functions()
}

// NewSimpleModule creates a Module with no sub-functions.
func NewSimpleModule(name string, run func(bool)) Module {
	return &simpleModule{name: name, run: run}
}

// NewModule creates a Module with named sub-functions.
func NewModule(name string, run func(bool), fns map[string]func(bool)) Module {
	return &fullModule{name: name, run: run, fns: fns}
}

type simpleModule struct {
	name string
	run  func(bool)
}

func (m *simpleModule) Name() string                     { return m.name }
func (m *simpleModule) Run(show bool)                    { m.run(show) }
func (m *simpleModule) Functions() map[string]func(bool) { return nil }

type fullModule struct {
	name string
	run  func(bool)
	fns  map[string]func(bool)
}

func (m *fullModule) Name() string                     { return m.name }
func (m *fullModule) Run(show bool)                    { m.run(show) }
func (m *fullModule) Functions() map[string]func(bool) { return m.fns }