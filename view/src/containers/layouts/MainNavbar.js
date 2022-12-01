import React from "react";
import classNames from "classnames";

const MainNavbar = ({ layout, stickyTop }) => {
  const classes = classNames(
    "navbar", "navbar-dark", "sticky-top", "bg-dark", "flex-md-nowrap", "p-0", "shadow"
  );

  return (
    <header className={classes}>
      <a className="navbar-brand col-md-3 col-lg-2 me-0 px-3" href="/">GUS</a>
    </header>
  );
};

export default MainNavbar;
