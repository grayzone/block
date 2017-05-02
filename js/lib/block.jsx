import React from "react";
import $ from "jquery";
import { Layer, Rect, Stage } from "react-konva";

class Block extends React.Component {
  constructor(props) {
    super(props);

    this.state = {
      color: Konva.Util.getRandomColor()
    };

    this.handleClick = this.handleClick.bind(this);
  }
  handleClick = e => {
    this.setState({
      color: Konva.Util.getRandomColor()
    });
    console.log(
      "click the block.x=",
      this.props.x,
      ":,y=",
      this.props.y,
      ",color=",
      this.state.color
    );
  };
  render() {
    return (
      <Rect
        x={this.props.x}
        y={this.props.y}
        width={50}
        height={50}
        shadowBlur={5}
        fill={this.state.color}
        onClick={this.handleClick}
      />
    );
  }
}

class Box extends React.Component {
  render() {
    const boxArray = [];
    for (let i = 0; i < 10; i++) {
      for (let j = 0; j < 10; j++) {
        boxArray.push(<Block x={i * 52} y={j * 52} />);
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
  render() {
    return (
      <Stage width={600} height={600}>
        <Box />
      </Stage>
    );
  }
}
