import React from "react";
import Container from 'react-bootstrap/Container';
import Row from 'react-bootstrap/Row';

import MainNavbar from "./MainNavbar";
import MainSidebar from "./MainSidebar";

const DefaultLayout = ({ children }) => (
  <div>
    <MainNavbar />
    <Container fluid>
      <Row>
        <MainSidebar />
        <main className="col-md-9 ms-sm-auto col-lg-10 px-md-4">
          <div className="justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
          {children}
          </div>
        </main>
      </Row>
    </Container>
  </div>
);

export default DefaultLayout;
