import './App.css';
import React from 'react';
import { Navigate, Route, createBrowserRouter, createRoutesFromElements, RouterProvider, useParams } from 'react-router-dom';
import DefaultLayout from './containers/layouts/DefaultLayout';
import Subjects from './containers/Subjects'
import 'bootstrap/dist/css/bootstrap.min.css';

function App() {
  const routes = [
    {
      path: "/",
      layout: DefaultLayout,
      element: () => <Navigate to="/subjects" />
    },
    {
      path: "/subjects/tree",
      layout: DefaultLayout,
      element: () => <Subjects/>
    },
    {
      path: "/subjects",
      layout: DefaultLayout,
      element: () => <Subjects/>
    },
    {
      path: "/subjects/:subjectId",
      layout: DefaultLayout,
      element: () => <Subjects {...useParams()}/>
    },
  ];
  const router = createBrowserRouter(
    createRoutesFromElements(
      routes.map((route, index) => {
        return (
          <Route
            key={index}
            path={route.path}
            loader={({ params }) => {
              console.log(params);
              return params;
            }}
            element={
              <route.layout>
                <route.element/>
              </route.layout>
            }
          />
        );
      })
    )
  )
  return <RouterProvider router={router} />;
}

export default App;
