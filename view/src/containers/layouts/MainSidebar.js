import React from "react";
import PropTypes from "prop-types";
import classNames from "classnames";
import { Col } from 'react-bootstrap';
import SidebarNavItems from "./SidebarNavItems";

class MainSidebar extends React.Component {
  render() {
    return (
      <nav id="sidebarMenu" className="col-md-3 col-lg-2 d-md-block bg-light sidebar collapse">
        <div className="position-sticky pt-3">
          <ul className="nav flex-column">
            <SidebarNavItems/>
          </ul>
        </div>        
        <hr/>
      </nav>
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
