import React from "react";
import ReactDOM from "react-dom";
import BlockBox from "./lib/block";

var blockDIV = document.getElementById("block");
if (blockDIV != null) {
  ReactDOM.render(<BlockBox />, blockDIV);
}
