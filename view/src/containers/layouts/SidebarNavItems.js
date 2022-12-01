import React from "react";
import Nav from 'react-bootstrap/Nav';
import SidebarNavItem from "./SidebarNavItem";

class SidebarNavItems extends React.Component {
  render() {
    // const { navItems: items } = this.state || [];
    const items = [
        {
            title: "Subjects",
            to: '/subjects'
        },
        {
            title: "Tree",
            to: '/subjects/tree'
        }
    ];
    return (
      <div>
        <Nav>
          {items.map((item, idx) => (
            <SidebarNavItem key={idx} item={item} />
          ))}
        </Nav>
      </div>
    )
  }
}

export default SidebarNavItems;
