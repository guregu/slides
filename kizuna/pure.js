var React = require("react");

var Message = React.createClass({
  mixins: [PureRenderMixin],　
  getInitialState: function () {
    return {
      liked: false
    };
  },
  render: function() {
    return (
      <li>
        {this.props.message} 
        <button disabled={this.state.liked}>Like</button>
      </li>
    );
  }
});

var MessageList = react.createClass({
  mixins: [PureRenderMixin],　
  getInitialState: function () {
    return {
      messages: Immutable.List()
    };
  },
  addMessage: function() {
    var msgs = this.state.messages;
    this.setState({
      messages: msgs.push("hello @ " + new Date)
    })
  },
  render: function() {
    return (
      <ul>
        <li><button onClick={this.addMessage} />Add message</button></li>
        {messages.forEach(function (msg, i) {
          return <Message message={msg} key={i} />
        })}
      </ul>
    );
  }
});

module.exports = Message;