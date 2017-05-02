import React from "react";
import $ from "jquery";
import { Layer, Rect, Stage } from "react-konva";

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

  render() {
    return (
      <Rect
        x={this.props.x}
        y={this.props.y}
        width={this.props.width}
        height={this.props.height}
        shadowBlur={this.state.shadowBlur}
        fill={this.props.color}
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
    const width = 50;
    const height = 50;
    const colors = ["white", "red", "blue", "yellow", "green", "purple"];
    var sizeX = width + 2;
    var sizeY = height + 2;
    for (let i = 0; i < 10; i++) {
      for (let j = 0; j < 10; j++) {
        let colorID = data[i][j];
        boxArray.push(
          <Block
            x={i * sizeX}
            y={j * sizeY}
            width={width}
            height={height}
            color={colors[colorID]}
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

class ActionButton extends React.Component {
  handleNewClick = e => {
    console.log("a new game.");
    this.props.onNewClick("this.initData()");
  };
  handleStartClick = e => {
    console.log("start the game.");
    this.props.onStartClick(e);
  };
  render() {
    return (
      <div>
        <button onClick={this.handleNewClick}>New</button>
        <button onClick={this.handleStartClick}>Start</button>
      </div>
    );
  }
}

export default class BlockBox extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      data: []
    };
  }
  getSeedData = () => {
    var url = "/new";
    $.ajax({
      url: url,
      dataType: "json",
      type: "GET",
      cache: false,
      async: false,
      success: data => {
        console.log("seed data:", data);
        this.setState({ data });
      },
      error: (xhr, status, err) => {
        console.error(url, status, err.toString());
      }
    });
  };

  handleNewButton = () => {
    console.log("handle new button:");
    this.getSeedData();
    /*
    this.setState({
      data: this.initData()
    });
    */
  };

  
  componentWillMount() {
    this.getSeedData();
  }
  

  handleStartButton = data => {
    console.log("handle start button.");
    this.setState({
      data
    });
  };

  render() {
    return (
      <div>
        <Stage width={600} height={600}>
          <Box data={this.state.data} />
        </Stage>
        <ActionButton
          onNewClick={this.handleNewButton}
          onStartClick={this.handleStartButton}
        />
      </div>
    );
  }
}
