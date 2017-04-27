import React from "react";
import { Button } from "antd";
import $ from "jquery";

export default class Block extends React.Component {
  handleNewClick = e => {
    var url = "/donew";
    $.ajax({
      url: url,
      dataType: "json",
      type: "GET",
      cache: false,
      async: false,
      success: data => {
        console.log("turn to the new game successfully:", data);
        window.location.href = data;
      },
      error: (xhr, status, err) => {
        console.error(url, status, err.toString());
      }
    });
  };
  handleNextClick = e => {
    var url = "/donext";
    $.ajax({
      url: url,
      dataType: "json",
      type: "GET",
      cache: false,
      async: false,
      success: data => {
        console.log("turn to the next game successfully:", data);
        window.location.href = data;
      },
      error: (xhr, status, err) => {
        console.error(url, status, err.toString());
      }
    });
  };
  render() {
    return (
      <div>
        This is a block game.<br />
        <Button onClick={this.handleNewClick}>New</Button>
        <Button onClick={this.handleNextClick}>Next</Button>
      </div>
    );
  }
}
