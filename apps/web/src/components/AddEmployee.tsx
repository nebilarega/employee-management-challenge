import React, { useState, useRef } from 'react';
import Swal from 'sweetalert2';
interface Employee {
  id: number;
  firstName: string;
  lastName: string;
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
  setIsAdding(isadding: boolean): void;
}

const AddEmployee: React.FC<Employees> = ({ employees, setIsAdding }) => {
  const [firstName, setFirstName] = useState('');
  const [lastName, setLastName] = useState('');
  const [email, setEmail] = useState('');
  const [phoneNo, setPhoneNo] = useState('');
  const [gender, setGender] = useState('');
  const [dateOfBirth, setDateOfBirth] = useState('');
  const [country, setCountry] = useState('');
  const [region, setRegion] = useState('');
  const [city, setCity] = useState('');
  const [subCity, setSubCity] = useState('');
  const [woreda, setWoreda] = useState('');
  const [zone, setZone] = useState('');
  const [kebele, setKebele] = useState('');
  const [houseNo, setHouseNo] = useState('');
  const [departmentId, setDepartmentId] = useState(0);

  const textInput = useRef<HTMLInputElement>(null);

  //   useEffect(() => {
  //     if (textInput != null) textInput.current.focus();
  //   }, []);

  const handleAdd = (e: React.FormEvent) => {
    e.preventDefault();
    if (
      !firstName ||
      !lastName ||
      !email ||
      !dateOfBirth ||
      !country ||
      !gender
    ) {
      Swal.fire({
        icon: 'error',
        title: 'Error!',
        text: 'All fields are required.',
        showConfirmButton: true,
      });
    }

    const id = employees.length + 1;
    const newEmployee = {
      id,
      firstName,
      lastName,
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
    employees.push(newEmployee);
    // setEmployees(employees);
    setIsAdding(false);

    Swal.fire({
      icon: 'success',
      title: 'Added!',
      text: `${firstName} ${lastName}'s data has been Added.`,
      showConfirmButton: false,
      timer: 1500,
    });
  };
  // max-width: 1200px;
  //   padding: 0 1rem;
  //   margin-left: auto;
  //   margin-right: auto;
  return (
    <div className="max-w-[1200px] pr-4 pl-4 ml-auto mr-auto">
      <form onSubmit={handleAdd}>
        <h1>Add Employee</h1>
        <label htmlFor="firstName">First Name</label>
        <input
          id="firstName"
          type="text"
          ref={textInput}
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
        <label htmlFor="gender">Gender</label>
        <input
          id="gender"
          type="number"
          name="gender"
          value={gender}
          onChange={(e) => setGender(e.target.value)}
        />
        <label htmlFor="dateOfBirth">Country</label>
        <input
          id="name"
          type="string"
          name="name"
          value={country}
          onChange={(e) => setCountry(e.target.value)}
        />
        <label htmlFor="dateOfBirth">Region</label>
        <input
          id="name"
          type="string"
          name="name"
          value={region}
          onChange={(e) => setRegion(e.target.value)}
        />
        <label htmlFor="dateOfBirth">Date of Birth</label>
        <input
          id="dateOfBirth"
          type="date"
          name="dateOfBirth"
          value={dateOfBirth}
          onChange={(e) => setDateOfBirth(e.target.value)}
        />
        <label htmlFor="dateOfBirth">Department ID</label>
        <input
          id="Department ID"
          type="number"
          name="Department ID"
          value={departmentId}
          onChange={(e) => setDepartmentId(parseInt(e.target.value))}
        />
        <div className="mt-[30px] flex justify-evenly">
          <input className="pr-4" type="submit" value="Add" />
          <input
            className="ml-[12px] bg-transparent border border-solid border-[#cdcdcd]"
            type="button"
            value="Cancel"
            onClick={() => setIsAdding(false)}
          />
        </div>
      </form>
    </div>
  );
};

export default AddEmployee;
