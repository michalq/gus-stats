import React from "react";
import PropTypes from "prop-types";
import classNames from "classnames";
import Col from 'react-bootstrap/Col';

import SidebarMainNavbar from "./SidebarMainNavbar";
import SidebarNavItems from "./SidebarNavItems";

class MainSidebar extends React.Component {
  render() {
    const classes = classNames(
      "main-sidebar",
      "px-0",
      "col-12",
      "open"
    );

    return (
      <Col
        tag="aside"
        className={classes}
        lg={{ size: 2 }}
        md={{ size: 3 }}
      >
        <SidebarMainNavbar hideLogoText={false} />
        <SidebarNavItems />
      </Col>
    );
  }
}

MainSidebar.propTypes = {
  /**
   * Whether to hide the logo text, or not.
   */
  hideLogoText: PropTypes.bool
};

MainSidebar.defaultProps = {
  hideLogoText: false
};

export default MainSidebar;
