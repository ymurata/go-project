import React, { ReactElement } from "react";

type MainTemplateProps = {
  children: Array<ReactElement>;
};

const MainTemplate: React.FC<MainTemplateProps> = ({ children }) => {
  const header = children.find((child) => child.key === "header");
  const main = children.find((child) => child.key === "main");
  return (
    <>
      {header}
      {main}
    </>
  );
};

export default MainTemplate;
