import React from "react";
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
      <ul className="nav nav-pills flex-column mb-auto">
        {items.map((item, idx) => (
          <SidebarNavItem key={idx} item={item} />
        ))}
      </ul>
    )
  }
}

export default SidebarNavItems;
