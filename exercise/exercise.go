package exercise

// Exercise is the interface acting as contract for all
// exercises
type Exercise interface {
	// Init is the method to initialize exercise
	Init()

	// Run, is used to run the exercise
	Run()
}
