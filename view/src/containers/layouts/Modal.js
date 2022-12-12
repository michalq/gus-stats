import React, { Component } from 'react';
 
export default class Modal extends Component {
  constructor(params) {
    super(params);
    this.id = params.id || "exampleModal";
    this.title = params.title || "Default title"
    this.children = params.children;
  }
  render() {
    return (
    <div className="modal fade" id={ this.id } tabindex="-1" aria-labelledby={ this.id + "Label" } aria-hidden="true">
      <div className="modal-dialog modal-fullscreen">
      <div className="modal-content">
        <div className="modal-header">
        <h5 className="modal-title" id={ this.id + "Label" }>{ this.title }</h5>
        <button type="button" className="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
        </div>
        <div className="modal-body">
        { this.children }
        </div>
        <div class="modal-footer">
        <button type="button" className="btn btn-secondary" data-bs-dismiss="modal">Close</button>
        </div>
      </div>
      </div>
    </div>
    )
    }
}