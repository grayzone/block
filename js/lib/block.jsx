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
    /*
    this.setState({
      shadowBlur: 0
    });
    */
    this.props.onChange({ x: this.props.indexX, y: 10-this.props.indexY });
    console.log(
      "the block state is:",
      this.state,
      ",x=",
      this.props.x,
      ",y=",
      this.props.y
    );
  };

  componentWillMount() {}

  render() {
    const colors = ["white", "red", "blue", "yellow", "green", "purple"];
    const colorID = this.props.colorID;
    return (
      <Rect
        x={this.props.x}
        y={this.props.y}
        indexX={this.props.indexX}
        indexY={this.props.indexY}
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
    this.props.onChange(data);
  };
  render() {
    const boxArray = [];
    const data = this.props.data;
    const width = this.props.width;
    const height = this.props.height;

    var sizeX = width + 2;
    var sizeY = height + 2;
    for (let i = 0; i < 10; i++) {
      for (let j = 10; j > 0; j--) {
        let colorID = data[i][10 - j];
        boxArray.push(
          <Block
            x={i * sizeX}
            y={j * sizeY}
            indexX={i}
            indexY={j}
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

class ResultBox extends React.Component {
  constructor(props) {
    super(props);
  }

  render() {
    const input = this.props.data;
    console.log("result box input:", input);
    const boxArray = [];
    for (let i = 0; i < input.length; i++) {
      boxArray.push(
        <Col span={2}>
          <Stage width={150} height={150}>
            <Box data={input[i].Data} width={10} height={10} />
          </Stage>
        </Col>
      );
    }
    return <div>{boxArray}</div>;
  }
}

export default class BlockBox extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      data: [],
      result: [],
      step: []
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

  handleBoxChange = point => {
    console.log("changed point:", point);
    var url = "/remove";
    $.ajax({
      url: url,
      dataType: "json",
      type: "POST",
      cache: false,
      async: false,
      data: {
        x: point.x,
        y: point.y,
        data:this.state.result.join(),
      },
      success: data => {
        console.log("remove data:", data);
        this.setState({
          result: data.Data
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

  handleStepClick = () => {
    var url = "/step";
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
        this.setState({ step: data });
      },
      error: (xhr, status, err) => {
        console.error(url, status, err.toString());
      }
    });
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
              <Box data={this.state.data} width={25} height={25} />
            </Stage>
          </Col>
          <Col span={5}>
            <Stage width={300} height={300}>
              <Box
                data={this.state.result}
                width={25}
                height={25}
                onChange={this.handleBoxChange}
              />
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
            <Button type="primary" onClick={this.handleStepClick}>Step</Button>
          </Col>
          <Col span={2}>
            <Button type="primary" onClick={this.handleStartClick}>
              Start
            </Button>
          </Col>
        </Row>

        <Row>
          <ResultBox data={this.state.step} />
        </Row>
      </div>
    );
  }
}
