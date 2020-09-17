import React from "react";
import { Container } from "@material-ui/core";

import MainTemplate from "../Templates/MainTemplate";
import MainHeader from "../Organisms/MainHeader";

const Home: React.SFC = () => (
  <MainTemplate>
    <MainHeader key="header" />
    <Container key="main">ホーム</Container>
  </MainTemplate>
);

export default Home;
