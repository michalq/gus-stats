import React from "react";
import Container from 'react-bootstrap/Container';
import Row from 'react-bootstrap/Row';
import Col from 'react-bootstrap/Col';

import MainNavbar from "./MainNavbar";
import MainSidebar from "./MainSidebar";
import MainFooter from "./MainFooter";

const DefaultLayout = ({ children }) => (
  <Container fluid>
    <Row>
      <MainSidebar />
      <Col
        className="main-content"
        lg={{ size: 10, offset: 2 }}
        md={{ size: 9, offset: 3 }}
        sm="12"
        tag="main"
      >
        <MainNavbar />
        {children}
        <MainFooter />
      </Col>
    </Row>
  </Container>
);

export default DefaultLayout;
