import React, { Component } from 'react';

export default class SubjectVariables extends Component {
  constructor(props) {
    super(props);

    this.state = {
      variables: [],
    }
  }

  componentDidMount () {
    let url = 'http://localhost:3030/subjects/' + this.props.subjectId + '/variables';

    fetch(url).then(data => data.json())
      .then(data => this.setState({
        variables: data.data.variables
      }));
    }

  render() {
    console.log();
    const variables = this.state.variables.map((variable, i) => (
      <div className="mb-3 form-check">
        <input type="checkbox" className="form-check-input" id={ "variable-" + variable.id }/>
        <label className="form-check-label" for={ "variable-" + variable.id }>{ variable.name }</label>
      </div>
    ));
    return (
      <div>
        <div className='row'>
          <div class="col-sm-6">
            <div className='card'>
              <div className="card-header">
                Variables
              </div>
              <div className='card-body'>
                { variables }
              </div>
            </div>
          </div>
          <div class="col-sm-6">
            <div className='card'>
              <div className="card-header">
                Units
              </div>
              <div className='card-body'>
                <div className="mb-3 form-check">
                  <input type="checkbox" className="form-check-input" id={ "unit-1" }/>
                  <label className="form-check-label" for={ "unit-1" }>Polska</label>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    )
  }
}