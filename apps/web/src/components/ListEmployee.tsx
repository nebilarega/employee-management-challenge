import React from 'react';
import { GET_EMPLOYEES } from '../graphql/Queries';

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

  handleEdit(employee: number): void;
  handleDelete(editing: number): void;
}
const ListEmployee: React.FC<Employees> = ({
  employees,
  handleEdit,
  handleDelete,
}) => {
  return (
    <div className="md: w-full">
      <table>
        <thead>
          <tr className="odd: bg-[#f8f8f8]">
            <th>No.</th>
            <th>First Name</th>
            <th>Last Name</th>
            <th>Email</th>
            <th>Salary</th>
            <th>Date</th>
            <th colSpan={2} className="text-center">
              Actions
            </th>
          </tr>
        </thead>
        <tbody>
          {employees.length > 0 ? (
            employees.map((employee, i) => (
              <tr key={employee.id}>
                <td>{i + 1}</td>
                <td>{employee.firstName}</td>
                <td>{employee.lastName}</td>
                <td>{employee.email}</td>
                <td>{employee.phoneNo}</td>
                <td>{employee.dateOfBirth} </td>
                <td className="text-right">
                  <button
                    onClick={() => handleEdit(employee.id)}
                    className="text-[100%] leading-5 border border-solid"
                  >
                    Edit
                  </button>
                </td>
                <td className="text-left">
                  <button
                    onClick={() => handleDelete(employee.id)}
                    className="text-[100%] leading-5 border border-solid"
                  >
                    Delete
                  </button>
                </td>
              </tr>
            ))
          ) : (
            <tr>
              <td colSpan={7}>No Employees</td>
            </tr>
          )}
        </tbody>
      </table>
    </div>
  );
};

export default ListEmployee;
