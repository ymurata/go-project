import React from "react";
import { AppBar, Toolbar, IconButton } from "@material-ui/core";
import MenuIcon from "@material-ui/icons/Menu";

const MainHeader = () => (
  <AppBar position="static">
    <Toolbar>
      <IconButton edge="start" aria-label="menu">
        <MenuIcon />
      </IconButton>
    </Toolbar>
  </AppBar>
);

export default MainHeader;
