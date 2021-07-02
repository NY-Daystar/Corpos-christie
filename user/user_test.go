package user

import (
	"testing"
)

// For testing
// $ cd user
// $ go test -v

// Create user single with no children and check parts
func TestUserSingleOnlyIncome(t *testing.T) {
	var partsRef float64 = 0.

	var user *User = new(User)
	user.CalculateParts()
	t.Logf("User reference %+v", user)

	// Testing parts
	if partsRef != user.Parts {
		t.Errorf("Expected that the Parts \n%f\n should be equal to \n%v", partsRef, user.Parts)
	}
}

// Create user in couple with no children and check parts
func TestUserInCoupleNoChildren(t *testing.T) {
	var partsRef float64 = 2.

	var user *User = new(User)
	user.IsInCouple = true
	user.CalculateParts()
	t.Logf("User reference %+v", user)

	// Testing parts
	if partsRef != user.Parts {
		t.Errorf("Expected that the Parts \n%f\n should be equal to \n%v", partsRef, user.Parts)
	}

}

//Create user in couple with 3 children and check parts
func TestUserInCoupleWith3Children(t *testing.T) {
	var partsRef float64 = 3.5

	var user *User = new(User)
	user.IsInCouple = true
	user.Children = 3
	user.CalculateParts()
	t.Logf("User reference %+v", user)

	// Testing parts
	if partsRef != user.Parts {
		t.Errorf("Expected that the Parts \n%f\n should be equal to \n%v", partsRef, user.Parts)
	}
}
