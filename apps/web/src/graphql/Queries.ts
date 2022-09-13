import { gql } from '@apollo/client';

export const GET_EMPLOYEES = gql`
  query employees($searchTerm: String!, $page: PaginationInput!) {
    employees {
      EmployeeConnection
    }
  }
`;
export const GET_EMPLOYEE = gql`
  query employee($id: ID!) {
    employee {
      Employee
    }
  }
`;
