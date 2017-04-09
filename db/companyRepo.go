package db

import (
	"github.com/redright/shuttlebus/domain"
)

type CompanyRepo struct {
}

func (r *CompanyRepo) CreateCompany(company *domain.Company) {
	company.ID = GenerateID()
	Execute("INSERT INTO company (ID,Name) VALUES (?,?)", company.ID, company.Name)
}

func (r *CompanyRepo) AddClientToCompany(company *domain.Company, client *domain.Client) {
	Execute("INSERT INTO CompanyClient (CompanyID, ClientID) VALUES (?,?)", company.ID, client.ID)
}

func (r *CompanyRepo) CreateDriver(driver *domain.CompanyDriver) {
	driver.ID = GenerateID()
	if driver.UserID == "" {
		var userRepo UserRepo
		userRepo.CreateUser(&driver.User)
		driver.UserID = driver.User.ID
	}
	if driver.DriverID == "" {
		driver.Driver.ID = GenerateID()
		driver.DriverID = driver.Driver.ID
		Execute("INSERT INTO Driver (ID,UserID,LicenceNumber,LicenceTakenDate) VALUES(?,?,?,?)", driver.ID, driver.UserID, driver.LicenceNumber, driver.LicenceTakenDate)
	}
	Execute("INSERT INTO CompanyDriver (ID,CompanyID,DriverID,HireDate,Status) VALUES(?,?,?,?,?)", driver.ID, driver.CompanyID, driver.DriverID, driver.HireDate, driver.Status)
}

func (r *CompanyRepo) CreateVehicle(vehicle *domain.CompanyVehicle) {
	vehicle.ID = GenerateID()
	if vehicle.Vehicle.ID == "" {
		vehicle.VehicleID = GenerateID()
		vehicle.Vehicle.ID = vehicle.VehicleID
		Execute("insert into vehicle (ID,PlateNo,Capacity,BrandCode,ModelCode,ModelYear) VALUES(?,?,?,?,?,?)",
			vehicle.VehicleID, vehicle.PlateNo, vehicle.Capacity, vehicle.BrandCode, vehicle.ModelCode, vehicle.ModelYear)
	}
	Execute("INSERT INTO CompanyVehicle (ID,CompanyID,VehicleID) VALUES (?,?,?)", vehicle.ID, vehicle.CompanyID, vehicle.VehicleID)
}
