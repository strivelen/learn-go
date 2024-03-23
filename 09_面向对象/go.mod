module example.com/test-oop

go 1.22.0

replace example.com/student => ./student

require (
	example.com/person v0.0.0-00010101000000-000000000000
	example.com/student v0.0.0-00010101000000-000000000000
)

replace example.com/person => ./person
