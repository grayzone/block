import React from "react";
import $ from "jquery";
import { Layer, Rect, Stage } from "react-konva";
import { Button, Row, Col } from "antd";

class Block extends React.Component {
  constructor(props) {
    super(props);

    this.state = {
      shadowBlur: 5
    };

    this.handleClick = this.handleClick.bind(this);
  }
  handleClick = e => {
    console.log("click the block,", this.props);
    this.setState({
      shadowBlur: 0
    });
    console.log(
      "the block state is:",
      this.state,
      ",x=",
      this.props.x,
      ",y=",
      this.props.y
    );
  };

  componentWillMount() {
    /*
    if (this.props.colorID != 0) {
      this.setState({
        shadowBlur: 5
      });
    }
    */
  }

  render() {
    const colors = ["white", "red", "blue", "yellow", "green", "purple"];
    const colorID = this.props.colorID;
    return (
      <Rect
        x={this.props.x}
        y={this.props.y}
        width={this.props.width}
        height={this.props.height}
        shadowBlur={this.state.shadowBlur}
        fill={colors[colorID]}
        onClick={this.handleClick}
      />
    );
  }
}

class Box extends React.Component {
  handleBlockChange = data => {
    console.log("block changes:", data);
  };
  render() {
    const boxArray = [];
    const data = this.props.data;
    const width = 25;
    const height = 25;

    var sizeX = width + 2;
    var sizeY = height + 2;
    for (let i = 0; i < 10; i++) {
      for (let j = 10; j > 0; j--) {
        let colorID = data[i][10 - j];
        boxArray.push(
          <Block
            x={i * sizeX}
            y={j * sizeY}
            width={width}
            height={height}
            colorID={colorID}
            onChange={this.handleBlockChange}
          />
        );
      }
    }
    return (
      <Layer>
        {boxArray}
      </Layer>
    );
  }
}

export default class BlockBox extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      data: [],
      result: []
    };
  }
  getSeedData = () => {
    this.getData("new");
  };

  getTestData = () => {
    this.getData("test");
  };

  getData = action => {
    var url = "/" + action;
    $.ajax({
      url: url,
      dataType: "json",
      type: "GET",
      cache: false,
      async: false,
      success: data => {
        console.log("seed data:", data);
        this.setState({
          data: data,
          result: data
        });
      },
      error: (xhr, status, err) => {
        console.error(url, status, err.toString());
      }
    });
  };

  dropdata = () => {
    var url = "/drop";
    $.ajax({
      url: url,
      dataType: "json",
      type: "POST",
      cache: false,
      async: false,
      data: {
        data: this.state.data.join()
      },
      success: data => {
        console.log("seed data:", data);
        this.setState({ result: data });
      },
      error: (xhr, status, err) => {
        console.error(url, status, err.toString());
      }
    });
  };

  handleNewClick = () => {
    console.log("handle new button:");
    this.getSeedData();
  };

  handleTestClick = () => {
    this.getTestData();
  };

  componentWillMount() {
    //  this.getSeedData();
    this.getTestData();
  }

  handleStartClick = data => {
    this.dropdata();
    console.log("handle start button.");
  };

  render() {
    return (
      <div>
        <Row gutter={16}>
          <Col span={5}>
            <Stage width={300} height={300}>
              <Box data={this.state.data} />
            </Stage>
          </Col>
          <Col span={5}>
            <Stage width={300} height={300}>
              <Box data={this.state.result} />
            </Stage>
          </Col>
        </Row>

        <Row gutter={8}>
          <Col span={2}>
            <Button type="primary" onClick={this.handleNewClick}>New</Button>
          </Col>
          <Col span={2}>
            <Button type="primary" onClick={this.handleTestClick}>Test</Button>
          </Col>
          <Col span={2}>
            <Button type="primary" onClick={this.handleStartClick}>
              Start
            </Button>
          </Col>
        </Row>
      </div>
    );
  }
}
