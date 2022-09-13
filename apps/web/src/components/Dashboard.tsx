import { useState } from 'react';
import Swal from 'sweetalert2';

import Header from './Header';
import AddEmployee from './AddEmployee';
import UpdateEmployee from './UpdateEmployee';
import ListEmployee from './ListEmployee';

// $firstName: String!;
// $lastName: String!;
// $gender: String!;
// $phoneNo: String!;
// $email: String!;
// $dateOfBirth: Time!;
// $country: String!;
// $region: String;
// $city: String;
// $subCity: String;
// $woreda: String;
// $zone: String;
// $kebele: String;
// $houseNo: String;
// $departmentId: Int!;
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
const Dashboard = () => {
  const emptyEmployee = {
    id: 0,
    firstName: '',
    lastName: '',
    gender: '',
    phoneNo: '',
    email: '',
    dateOfBirth: '',
    country: '',
    region: '',
    city: '',
    subCity: '',
    woreda: '',
    zone: '',
    kebele: '',
    houseNo: '',
    departmentId: 0,
  };
  const [employees, setEmployees] = useState<Employee[]>([emptyEmployee]);
  const [selectedEmployee, setSelectedEmployee] =
    useState<Employee>(emptyEmployee);
  const [isAdding, setIsAdding] = useState<boolean>(false);
  const [isEditing, setIsEditing] = useState(false);

  const handleEdit = (id: number) => {
    const employee = employees.filter((employee) => employee.id === id);

    setSelectedEmployee(employee[0]);
    setIsEditing(true);
  };

  const handleDelete = (id: number) => {
    Swal.fire({
      icon: 'warning',
      title: 'Are you sure?',
      text: "You won't be able to revert this!",
      showCancelButton: true,
      confirmButtonText: 'Yes, delete it!',
      cancelButtonText: 'No, cancel!',
    }).then((result) => {
      if (result.value) {
        const [employee] = employees.filter((employee) => employee.id === id);

        Swal.fire({
          icon: 'success',
          title: 'Deleted!',
          text: `${employee.firstName} ${employee.lastName}'s data has been deleted.`,
          showConfirmButton: false,
          timer: 1500,
        });

        setEmployees(employees.filter((employee) => employee.id !== id));
      }
    });
  };

  return (
    <div className="max-w-[1200px] pr-4 pl-4 ml-auto mr-auto">
      {/* List */}
      {!isAdding && !isEditing && (
        <>
          <Header setIsAdding={setIsAdding} />
          <ListEmployee
            employees={employees}
            handleEdit={handleEdit}
            handleDelete={handleDelete}
          />
        </>
      )}
      {/* Add */}
      {isAdding && (
        <AddEmployee employees={employees} setIsAdding={setIsAdding} />
      )}
      {/* Edit */}
      {isEditing && (
        <UpdateEmployee
          employees={employees}
          selectedEmployee={selectedEmployee}
          setEmployees={setEmployees}
          setIsEditing={setIsEditing}
        />
      )}
    </div>
  );
};

export default Dashboard;
