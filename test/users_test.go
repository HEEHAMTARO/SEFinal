package unit

import (
	"fmt"
	"testing"
	"time"

	"github.com/asaskevich/govalidator"
	."github.com/onsi/gomega"

	"example.com/Sefinal/entity"
)

func TestUnit(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run("StudentID is required", func(t *testing.T) {
		users := entity.User{
			StudentID: "",
			FirstName: "Unit",
			LastName:  "test",
			Email:     "test@gmail.com",
			Phone:     "0800000000",
			Profile:   "",
			Password:  "",
			BirthDay:  time.Now(),
			LinkedIn:  "https://www.linkedin.com/company/ilink/",
			GenderID:  1,
		}

		ok, err := govalidator.ValidateStruct(users)

		g.Expect(ok).NotTo(BeTrue())

		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("StudentID is required"))
	})

	t.Run("StudentID is valid", func(t *testing.T){
		users := entity.User{
			StudentID: "B65113266",
			FirstName: "Unit",
			LastName:  "test",
			Email:     "test@gmail.com",
			Phone:     "0800000000",
			Profile:   "",
			Password:  "",
			BirthDay:  time.Now(),
			LinkedIn:  "https://www.linkedin.com/company/ilink/",
			GenderID:  1,
		}

		ok, err := govalidator.ValidateStruct(users)

		g.Expect(ok).NotTo(BeTrue())

		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal(fmt.Sprintf("StudentID: %s does not validate as matches(^[BMD]\\d{7}$)", users.StudentID)))
	})
}

func TestFirstName(t *testing.T){
	g := NewGomegaWithT(t)

	t.Run("FirstName is required", func(t *testing.T){
		users := entity.User{
			StudentID: "B6511326",
			FirstName: "",
			LastName:  "test",
			Email:     "test@gmail.com",
			Phone:     "0800000000",
			Profile:   "",
			Password:  "",
			BirthDay:  time.Now(),
			LinkedIn:  "https://www.linkedin.com/company/ilink/",
			GenderID:  1,
		}

		ok, err := govalidator.ValidateStruct(users)

		g.Expect(ok).NotTo(BeTrue())

		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("FirstName is required"))
	})
}

func TestUsersStudentID(t *testing.T) {
	g := NewGomegaWithT(t)
	t.Run("UsersID is required", func(t *testing.T) {
		users := entity.User{
			StudentID: "",
			FirstName: "Unit",
			LastName:  "test",
			Email:     "test@gmail.com",
			Phone:     "0800000000",
			Profile:   "",
			Password:  "",
			BirthDay:  time.Now(),
			LinkedIn:  "https://www.linkedin.com/company/ilink/",
			GenderID:  1,
		}

		ok, err := govalidator.ValidateStruct(users)

		g.Expect(ok).NotTo(BeTrue())

		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("StudentID is required"))

		t.Run("StudentID is valid", func(t *testing.T) {
			users := entity.User{
				StudentID: "B65113266",
				FirstName: "Unit",
				LastName:  "test",
				Email:     "test@gmail.com",
				Phone:     "0800000000",
				Profile:   "",
				Password:  "",
				BirthDay:  time.Now(),
				LinkedIn:  "https://www.linkedin.com/company/ilink/",
				GenderID:  1,
			}

			ok, err := govalidator.ValidateStruct(users)

			g.Expect(ok).NotTo(BeTrue())

			g.Expect(err).NotTo(BeNil())

			g.Expect(err.Error()).To(Equal(fmt.Sprintf("StudentID: %s does not validate as matches(^[BMD]\\d{7}$)", users.StudentID)))
		})
	})
}

func TestEmail(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run("Email is required", func(t *testing.T) {
		users := entity.User{
			StudentID: "B6511326",
			FirstName: "Unit",
			LastName:  "test",
			Email:     "",
			Phone:     "0800000000",
			Profile:   "",
			Password:  "",
			BirthDay:  time.Now(),
			LinkedIn:  "https://www.linkedin.com/company/ilink/",
			GenderID:  1,
		}

		ok, err := govalidator.ValidateStruct(users)

		g.Expect(ok).NotTo(BeTrue())

		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("Email is required"))

		t.Run("Email is validate", func(t *testing.T) {
			users := entity.User{
				StudentID: "B6511326",
				FirstName: "Unit",
				LastName:  "test",
				Email:     "testgmail.com",
				Phone:     "0800000000",
				Profile:   "",
				Password:  "",
				BirthDay:  time.Now(),
				LinkedIn:  "https://www.linkedin.com/company/ilink/",
				GenderID:  1,
			}

			ok, err := govalidator.ValidateStruct(users)

			g.Expect(ok).NotTo(BeTrue())

			g.Expect(err).NotTo(BeNil())

			g.Expect(err.Error()).To(Equal("Email is invalid"))
		})
	})
}

func TestPhone(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run("Phone is required", func(t *testing.T) {
		users := entity.User{
			StudentID: "B6511326",
			FirstName: "Unit",
			LastName:  "test",
			Email:     "test@gmail.com",
			Phone:     "",
			Profile:   "",
			Password:  "",
			BirthDay:  time.Now(),
			LinkedIn:  "https://www.linkedin.com/company/ilink/",
			GenderID:  1,
		}

		ok, err := govalidator.ValidateStruct(users)

		g.Expect(ok).NotTo(BeTrue())

		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("Phone is required"))

		t.Run("Phone is validate", func(t *testing.T) {
			users := entity.User{
				StudentID: "B6511326",
				FirstName: "Unit",
				LastName:  "test",
				Email:     "test@gmail.com",
				Phone:     "08000000001",
				Profile:   "",
				Password:  "",
				BirthDay:  time.Now(),
				LinkedIn:  "https://www.linkedin.com/company/ilink/",
				GenderID:  1,
			}

			ok, err := govalidator.ValidateStruct(users)

			g.Expect(ok).NotTo(BeTrue())

			g.Expect(err).NotTo(BeNil())

			g.Expect(err.Error()).To(Equal(fmt.Sprintf("Phone: %s does not validate as stringlength(10|10)", users.Phone)))
		})
	})
}

func TestUrl(t *testing.T) {

	g := NewGomegaWithT(t)

	t.Run("URL is required", func(t *testing.T) {
		users := entity.User{
			StudentID: "B6511326",
			FirstName: "Unit",
			LastName:  "test",
			Email:     "test@gmail.com",
			Phone:     "0800000000",
			Profile:   "",
			Password:  "",
			BirthDay:  time.Now(),
			LinkedIn:  "",
			GenderID:  1,
		}

		ok, err := govalidator.ValidateStruct(users)

		g.Expect(ok).NotTo(BeTrue())

		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("LinkedIn is required"))

		t.Run("Url is valid", func(t *testing.T) {
			users := entity.User{
				StudentID: "B6511326",
				FirstName: "Unit",
				LastName:  "test",
				Email:     "test@gmail.com",
				Phone:     "0800000000",
				Profile:   "",
				Password:  "",
				BirthDay:  time.Now(),
				LinkedIn:  "//www.linkedin.com/company/ilink/",
				GenderID:  1,
			}

			ok, err := govalidator.ValidateStruct(users)

			g.Expect(ok).NotTo(BeTrue())

			g.Expect(err).NotTo(BeNil())

			g.Expect(err.Error()).To(Equal("Url LinkedIn is invalid"))
		})
	})
}

func TestUserID(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run("StudentID is required", func(t *testing.T) {
		users := entity.User{
			StudentID: "",
			FirstName: "Unit",
			LastName:  "test",
			Email:     "test@gmail.com",
			Phone:     "0800000000",
			Profile:   "",
			Password:  "",
			BirthDay:  time.Now(),
			LinkedIn:  "https://www.linkedin.com/company/ilink/",
			GenderID:  1,
		}

		ok, err := govalidator.ValidateStruct(users)

		g.Expect(ok).NotTo(BeTrue())

		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("StudentID is required"))
	})

	t.Run("StudentID is valid", func(t *testing.T) {
		users := entity.User{
			StudentID: "B65113266",
			FirstName: "Unit",
			LastName:  "test",
			Email:     "test@gmail.com",
			Phone:     "0800000000",
			Profile:   "",
			Password:  "",
			BirthDay:  time.Now(),
			LinkedIn:  "https://www.linkedin.com/company/ilink/",
			GenderID:  1,
		}

		ok, err := govalidator.ValidateStruct(users)

		g.Expect(ok).NotTo(BeTrue())

		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal(fmt.Sprintf("StudentID: %s does not validate as matches(^[BMD]\\d{7}$)", users.StudentID)))
	})
}

func TestUsersEmail(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run("Email is required", func(t *testing.T) {
		users := entity.User{
			StudentID: "B6511326",
			FirstName: "Unit",
			LastName:  "test",
			Email:     "",
			Phone:     "0800000000",
			Profile:   "",
			Password:  "",
			BirthDay:  time.Now(),
			LinkedIn:  "https://www.linkedin.com/company/ilink/",
			GenderID:  1,
		}

		ok, err := govalidator.ValidateStruct(users)

		g.Expect(ok).NotTo(BeTrue())

		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("Email is required"))
	})

	t.Run("Email is validation", func(t *testing.T) {
		users := entity.User{
			StudentID: "B6511326",
			FirstName: "Unit",
			LastName:  "test",
			Email:     "testgmail.com",
			Phone:     "0800000000",
			Profile:   "",
			Password:  "",
			BirthDay:  time.Now(),
			LinkedIn:  "https://www.linkedin.com/company/ilink/",
			GenderID:  1,
		}

		ok, err := govalidator.ValidateStruct(users)

		g.Expect(ok).NotTo(BeTrue())

		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("Email is invalid"))
	})
}

func TestUsersPayment(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run("Phone is required", func(t *testing.T) {
		users := entity.User{
			StudentID: "B6511326",
			FirstName: "Unit",
			LastName:  "test",
			Email:     "test@gmail.com",
			Phone:     "",
			Profile:   "",
			Password:  "",
			BirthDay:  time.Now(),
			LinkedIn:  "https://www.linkedin.com/company/ilink/",
			GenderID:  1,
		}

		ok, err := govalidator.ValidateStruct(users)

		g.Expect(ok).NotTo(BeTrue())

		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("Phone is required"))
	})

	t.Run("Phone is validate", func(t *testing.T) {
		users := entity.User{
			StudentID: "B6511326",
			FirstName: "Unit",
			LastName:  "test",
			Email:     "test@gmail.com",
			Phone:     "08000000001",
			Profile:   "",
			Password:  "",
			BirthDay:  time.Now(),
			LinkedIn:  "https://www.linkedin.com/company/ilink/",
			GenderID:  1,
		}

		ok, err := govalidator.ValidateStruct(users)

		g.Expect(ok).NotTo(BeTrue())

		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal(fmt.Sprintf("Phone: %s does not validate as stringlength(10|10)", users.Phone)))
	})
}

// func TestStudentID(t *testing.T) {

// 	g := NewGomegaWithT(t)

// 	t.Run(`student_id is required`, func(t *testing.T) {
// 		user := entity.User{
// 			StudentID: "", // ผิดตรงนี้
// 			FirstName: "Unit",
// 			LastName:  "test",
// 			Email:     "test@gmail.com",
// 			Phone:     "0800000000",
// 			Profile:   "",
// 			Password:  "",
// 			BirthDay:  time.Now(),
// 			LinkedIn:  "https://www.linkedin.com/company/ilink/",
// 			GenderID:  1,
// 		}

// 		// ตรวจสอบด้วย govalidator
// 		ok, err := govalidator.ValidateStruct(user)

// 		// ok ต้องไม่เป็นค่า true แปลว่าต้องจับ error ได้
// 		g.Expect(ok).NotTo(BeTrue())
// 		// err ต้องไม่เป็นค่า nil แปลว่าต้องจับ error ได้
// 		g.Expect(err).NotTo(BeNil())

// 		// err.Error ต้องมี error message แสดงออกมา
// 		g.Expect(err.Error()).To(Equal("StudentID is required"))
// 	})

// t.Run(`student_id pattern is not true`, func(t *testing.T) {
// 	user := entity.User{
// 		StudentID: "K5000000", // ผิดตรงนี้
// 		FirstName: "unit",
// 		LastName:  "test",
// 		Email:     "test@gmail.com",
// 		Phone:     "0800000000",
// 		Profile:   "",
// 		Password:  "",
// 		BirthDay:  time.Now(),
// 		LinkedIn:  "https://www.linkedin.com/company/ilink/",
// 		GenderID:  1,
// 	}

// 	// ตรวจสอบด้วย govalidator
// 	ok, err := govalidator.ValidateStruct(user)

// 	// ok ต้องไม่เป็นค่า true แปลว่าต้องจับ error ได้
// 	g.Expect(ok).NotTo(BeTrue())
// 	// err ต้องไม่เป็นค่า nil แปลว่าต้องจับ error ได้
// 	g.Expect(err).NotTo(BeNil())

// 	// err.Error ต้องมี error message แสดงออกมา
// 	g.Expect(err.Error()).To(Equal(fmt.Sprintf("StudentID: %s does not validate as matches(^[BMD]\\d{7}$)", user.StudentID)))
// })
// }

func TestUser(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run(`StudentID is required`, func(t *testing.T) {
		user := entity.User{
			StudentID: ``,
			FirstName: "unit",
			LastName:  "test",
			Email:     "test@gmail.com",
			Phone:     "0800000000",
			Profile:   "",
			Password:  "",
			BirthDay:  time.Now(),
			LinkedIn:  "https://www.linkedin.com/company/ilink/",
			GenderID:  1,
		}

		ok, err := govalidator.ValidateStruct(user)

		g.Expect(ok).NotTo(BeTrue())

		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("StudentID is required"))

	})

	t.Run(`StudentID is invalid`, func(t *testing.T) {
		user := entity.User{
			StudentID: `B00000000`,
			FirstName: "unit",
			LastName:  "test",
			Email:     "test@gmail.com",
			Phone:     "0800000000",
			Profile:   "",
			Password:  "",
			BirthDay:  time.Now(),
			LinkedIn:  "https://www.linkedin.com/company/ilink/",
			GenderID:  1,
		}

		ok, err := govalidator.ValidateStruct(user)

		g.Expect(ok).NotTo(BeTrue())

		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal(fmt.Sprintf("StudentID: %s does not validate as matches(^[BMD]\\d{7}$)", user.StudentID)))

	})

}

func TestPhoneNumber(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run(`phone_number is required`, func(t *testing.T) {
		user := entity.User{
			StudentID: "B5000000",
			FirstName: "Unit",
			LastName:  "test",
			Email:     "test@gmail.com",
			Phone:     "", // ผิดตรงนี้
			Profile:   "",
			Password:  "",
			BirthDay:  time.Now(),
			LinkedIn:  "https://www.linkedin.com/company/ilink/",
			GenderID:  1,
		}
		// ตรวจสอบด้วย govalidator
		ok, err := govalidator.ValidateStruct(user)

		// ok ต้องไม่เป็นค่า true แปลว่าต้องจับ error ได้
		g.Expect(ok).NotTo(BeTrue())
		// err ต้องไม่เป็นค่า nil แปลว่าต้องจับ error ได้
		g.Expect(err).NotTo(BeNil())

		// err.Error ต้องมี error message แสดงออกมา
		g.Expect(err.Error()).To(Equal("Phone is required"))

	})

	t.Run(`phone_number check 10 digit`, func(t *testing.T) {
		user := entity.User{
			StudentID: "B5000000",
			FirstName: "Unit",
			LastName:  "test",
			Email:     "test@gmail.com",
			Phone:     "080800000000", // ผิดตรงนี้ มี 11 ตัว
			Profile:   "",
			Password:  "",
			BirthDay:  time.Now(),
			LinkedIn:  "https://www.linkedin.com/company/ilink/",
			GenderID:  1,
		}

		// ตรวจสอบด้วย govalidator
		ok, err := govalidator.ValidateStruct(user)

		// ok ต้องไม่เป็นค่า true แปลว่าต้องจับ error ได้
		g.Expect(ok).NotTo(BeTrue())
		// err ต้องไม่เป็นค่า nil แปลว่าต้องจับ error ได้
		g.Expect(err).NotTo(BeNil())

		// err.Error ต้องมี error message แสดงออกมา
		g.Expect(err.Error()).To(Equal(fmt.Sprintf("Phone: %s does not validate as stringlength(10|10)", user.Phone)))

	})
}

func TestValid(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run(`valid`, func(t *testing.T) {
		user := entity.User{
			StudentID: "M6601140",
			FirstName: "unit",
			LastName:  "test",
			Email:     "test@gmail.com",
			Phone:     "0800000000",
			Profile:   "",
			Password:  "",
			BirthDay:  time.Now(),
			LinkedIn:  "https://www.linkedin.com/company/ilink/",
			GenderID:  1,
		}

		// ตรวจสอบด้วย govalidator
		ok, err := govalidator.ValidateStruct(user)

		// ok ต้องไม่เป็นค่า true แปลว่าต้องจับ error ได้
		g.Expect(ok).NotTo(BeTrue())
		// err ต้องไม่เป็นค่า nil แปลว่าต้องจับ error ได้
		g.Expect(err).NotTo(BeNil())
	})
}

func TestTrue(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run("Valid", func(t *testing.T){
		users := entity.User{
			StudentID: "B6511326",
			FirstName: "Unit",
			LastName:  "test",
			Email:     "test@gmail.com",
			Phone:     "0800000000",
			Profile:   "",
			Password:  "",
			BirthDay:  time.Now(),
			LinkedIn:  "https://www.linkedin.com/company/ilink/",
			GenderID:  1,
		}

		ok, err := govalidator.ValidateStruct(users)

		g.Expect(ok).To(BeTrue())

		g.Expect(err).To(BeNil())

	})
}
