// eslint-disable-next-line @typescript-eslint/no-unused-vars
import { ApolloClient, InMemoryCache, ApolloProvider } from '@apollo/client';
import Dashboard from '../components/Dashboard';
export function App() {
  const client = new ApolloClient({
    uri: ' http://localhost:8080/query',
    cache: new InMemoryCache(),
  });

  return (
    <>
      <ApolloProvider client={client}>
        <Dashboard />
      </ApolloProvider>
      <div></div>
    </>
  );
}

export default App;
