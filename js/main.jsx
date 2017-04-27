import React from "react";
import ReactDOM from "react-dom";
import Block from "./lib/block";
import NewGame from "./lib/new";
import NextGame from "./lib/next";

var blockDIV = document.getElementById("block");
if (blockDIV != null) {
  ReactDOM.render(<Block />, blockDIV);
}

var newDIV = document.getElementById("new");
if (newDIV != null) {
  ReactDOM.render(<NewGame />, newDIV);
}

var nextDIV = document.getElementById("next");
if (nextDIV != null) {
  ReactDOM.render(<NextGame />, nextDIV);
}