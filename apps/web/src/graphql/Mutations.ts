import { gql } from '@apollo/client';

// firstName: String!;
// lastName: String!;
// gender: String!;
// phoneNo: String!;
// email: String!;
// dateOfBirth: Time!;
// country: String!;
// region: String;
// city: String;
// subCity: String;
// woreda: String;
// zone: String;
// kebele: String;
// houseNo: String;
// departmentId: Int!;

// id: ID!;
// firstName: String!;
// lastName: String!;
// gender: String!;
// phoneNo: String!;
// email: String!;
// dateOfBirth: Time!;
// country: String!;
// region: String;
// city: String;
// subCity: String;
// woreda: String;
// zone: String;
// kebele: String;
// houseNo: String;
// departmentId: Int!;

export const CREATE_EMPLOYEE = gql`
  mutation createEmployee(
    $firstName: String!
    $lastName: String!
    $gender: String!
    $phoneNo: String!
    $email: String!
    $dateOfBirth: Time!
    $country: String!
    $region: String
    $city: String
    $subCity: String
    $woreda: String
    $zone: String
    $kebele: String
    $houseNo: String
    $departmentId: Int!
  ) {
    createEmployee(
      firstName: $firstName
      lastName: $lastName
      gender: $gender
      phoneNo: $phoneNo
      email: $email
      dateOfBirth: $dateOfBirth
      country: $country
      region: $region
      city: $city
      subCity: $subCity
      woreda: $woreda
      zone: $zone
      kebele: $kebele
      houseNo: $houseNo
      departmentId: $departmentId
    ) {
      firstName
      lastName
      departmentId
    }
  }
`;

export const UPDATE_EMPLOYEE = gql`
  mutation updateEmployee(
    $id: ID!
    $firstName: String!
    $lastName: String!
    $gender: String!
    $phoneNo: String!
    $email: String!
    $dateOfBirth: Time!
    $country: String
    $region: String
    $city: String
    $subCity: String
    $woreda: String
    $zone: String
    $kebele: String
    $houseNo: String
    $departmentId: Int
  ) {
    updateEmployee(
      id: $id
      firstName: $firstName
      lastName: $lastName
      gender: $gender
      phoneNo: $phoneNo
      email: $email
      dateOfBirth: $dateOfBirth
      country: $country
      region: $region
      city: $city
      subCity: $subCity
      woreda: $woreda
      zone: $zone
      kebele: $kebele
      houseNo: $houseNo
      departmentId: $departmentId
    ) {
      firstName
      lastName
      departmentId
    }
  }
`;

export const DELETE_EMPLOYEE = gql`
  mutation deleteEmployee($id: ID!) {
    deleteEmployee(id: $id) {
      message
    }
  }
`;
