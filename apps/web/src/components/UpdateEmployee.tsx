import React, { useState } from 'react';
import Swal from 'sweetalert2';

interface Employee {
  id: number;
  firstName: string;
  lastName: string;
  gender: string;
  phoneNo: string;
  email: string;
  dateOfBirth: string;
  country: string;
  region: string;
  city: string;
  subCity: string;
  woreda: string;
  zone: string;
  kebele: string;
  houseNo: string;
  departmentId: number;
}
interface Employees {
  employees: Employee[];
  selectedEmployee: Employee;
  setEmployees(employee: Employee[]): void;
  setIsEditing(editing: boolean): void;
}
const UpdateEmployee: React.FC<Employees> = ({
  employees,
  selectedEmployee,
  setEmployees,
  setIsEditing,
}) => {
  const id = selectedEmployee.id;

  const [firstName, setFirstName] = useState(selectedEmployee.firstName);
  const [lastName, setLastName] = useState(selectedEmployee.lastName);
  const [email, setEmail] = useState(selectedEmployee.email);
  const [phoneNo, setPhoneNo] = useState(selectedEmployee.phoneNo);
  const [gender, setGender] = useState(selectedEmployee.gender);
  const [dateOfBirth, setDateOfBirth] = useState(selectedEmployee.dateOfBirth);
  const [country, setCountry] = useState(selectedEmployee.country);
  const [region, setRegion] = useState(selectedEmployee.region);
  const [city, setCity] = useState(selectedEmployee.city);
  const [subCity, setSubCity] = useState(selectedEmployee.subCity);
  const [woreda, setWoreda] = useState(selectedEmployee.woreda);
  const [zone, setZone] = useState(selectedEmployee.zone);
  const [kebele, setKebele] = useState(selectedEmployee.kebele);
  const [houseNo, setHouseNo] = useState(selectedEmployee.houseNo);
  const [departmentId, setDepartmentId] = useState(
    selectedEmployee.departmentId
  );

  const handleUpdate = (e: React.FormEvent) => {
    e.preventDefault();

    if (
      !firstName ||
      !lastName ||
      !email ||
      !phoneNo ||
      !dateOfBirth ||
      !gender ||
      !country
    ) {
      Swal.fire({
        icon: 'error',
        title: 'Error!',
        text: 'All fields are required.',
        showConfirmButton: true,
      });
    }

    const employee = {
      id,
      firstName,
      lastName,
      gender,
      phoneNo,
      email,
      dateOfBirth,
      country,
      region,
      city,
      subCity,
      woreda,
      zone,
      kebele,
      houseNo,
      departmentId,
    };

    for (let i = 0; i < employees.length; i++) {
      if (employees[i].id === id) {
        employees.splice(i, 1, employee);
        break;
      }
    }

    setEmployees(employees);
    setIsEditing(false);

    Swal.fire({
      icon: 'success',
      title: 'Updated!',
      text: `${employee.firstName} ${employee.lastName}'s data has been updated.`,
      showConfirmButton: false,
      timer: 1500,
    });
  };

  return (
    <div className="max-w-[1200px] pr-4 pl-4 ml-auto mr-auto">
      <form onSubmit={handleUpdate}>
        <h1>Edit Employee</h1>
        <label htmlFor="firstName">First Name</label>
        <input
          id="firstName"
          type="text"
          name="firstName"
          value={firstName}
          onChange={(e) => setFirstName(e.target.value)}
        />
        <label htmlFor="lastName">Last Name</label>
        <input
          id="lastName"
          type="text"
          name="lastName"
          value={lastName}
          onChange={(e) => setLastName(e.target.value)}
        />
        <label htmlFor="email">Email</label>
        <input
          id="email"
          type="email"
          name="email"
          value={email}
          onChange={(e) => setEmail(e.target.value)}
        />
        <label htmlFor="phoneNo">phoneNo</label>
        <input
          id="phoneNo"
          type="number"
          name="phoneNo"
          value={phoneNo}
          onChange={(e) => setPhoneNo(e.target.value)}
        />
        <label htmlFor="">Gender</label>
        <input
          id="gender"
          type="string"
          name="gender"
          value={dateOfBirth}
          onChange={(e) => setGender(e.target.value)}
        />
        <label htmlFor="date">Country</label>
        <input
          id="date"
          type="date"
          name="date"
          value={dateOfBirth}
          onChange={(e) => setCountry(e.target.value)}
        />
        <label htmlFor="date">Department ID</label>
        <input
          id="departmentID"
          type="number"
          name="date"
          value={dateOfBirth}
          onChange={(e) => setDepartmentId(parseInt(e.target.value))}
        />
        <div className="mt-[30px]">
          <input type="submit" value="Update" />
          <input
            className="ml-[12px] bg-transparent border border-solid border-[#cdcdcd]"
            type="button"
            value="Cancel"
            onClick={() => setIsEditing(false)}
          />
        </div>
      </form>
    </div>
  );
};

export default UpdateEmployee;
