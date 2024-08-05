package generics

import "testing"

func TestStack(t *testing.T) {
	t.Run("integer stack", func(t *testing.T) {
		stackOfInts := new(Stack[int])

		AssertTrue(t, stackOfInts.IsEmpty())

		stackOfInts.Push(123)
		AssertFalse(t, stackOfInts.IsEmpty())

		stackOfInts.Push(456)
		value, _ := stackOfInts.Pop()
		AssertEqual(t, value, 456)
		value, _ = stackOfInts.Pop()
		AssertEqual(t, value, 123)
		AssertTrue(t, stackOfInts.IsEmpty())

		stackOfInts.Push(1)
		stackOfInts.Push(2)
		firstNum, _ := stackOfInts.Pop()
		secondNum, _ := stackOfInts.Pop()
		AssertEqual(t, firstNum+secondNum, 3)
	})

	t.Run("string stack", func(t *testing.T) {
		stackOfStrings := new(Stack[string])

		AssertTrue(t, stackOfStrings.IsEmpty())

		stackOfStrings.Push("123")
		AssertFalse(t, stackOfStrings.IsEmpty())

		stackOfStrings.Push("456")
		value, _ := stackOfStrings.Pop()
		AssertEqual(t, value, "456")
		value, _ = stackOfStrings.Pop()
		AssertEqual(t, value, "123")
		AssertTrue(t, stackOfStrings.IsEmpty())
	})

	t.Run("float stack", func(t *testing.T) {
		stackOfFloat := new(Stack[float64])

		AssertTrue(t, stackOfFloat.IsEmpty())

		stackOfFloat.Push(12.3)
		AssertFalse(t, stackOfFloat.IsEmpty())

		stackOfFloat.Push(4.5)
		value, _ := stackOfFloat.Pop()
		AssertEqual(t, value, 4.5)

		value, _ = stackOfFloat.Pop()
		AssertEqual(t, value, 12.3)
		AssertTrue(t, stackOfFloat.IsEmpty())
	})
}
