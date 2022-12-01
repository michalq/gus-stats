import React from "react";

const SidebarNavItem = ({ item }) => (
  <li className="nav-item">
    <a className="nav-link" href={item.to}>
      <span data-feather="home"></span>
      {item.title}
    </a>
  </li>
);

export default SidebarNavItem;
