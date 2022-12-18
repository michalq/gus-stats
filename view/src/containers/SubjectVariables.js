import React, { Component } from 'react';
import LineChartComponent from './graphs/LineChart';
import Topo from './graphs/Topo';
import Modal from './layouts/Modal';

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
          <div className='col-sm-12'>
            <div class="card">
              <div className='card-body'>
                <button type="button" class="btn btn-primary" data-bs-toggle="modal" data-bs-target="#chart-regular">Render graph</button>
                <button type="button" class="btn btn-primary" data-bs-toggle="modal" data-bs-target="#chart-map-voievodeship">Voievodeship</button>
                <Modal id="chart-regular" title="Regular chart">
                  <LineChartComponent/>
                </Modal>
                <Modal id="chart-map-voievodeship" title="Voievodeship map">
                  <Topo/>
                </Modal>
              </div>
            </div>
          </div>
        </div>
        <div className='row'>
          <div class="col-sm-6">
            <div className='card'>
              <div className='card-body'>
                <h5 className="card-title">Variables</h5>
                { variables }
              </div>
            </div>
          </div>
          <div class="col-sm-6">
            <div className='card'>
              <div className='card-body'>
                <h5 className="card-title">Units</h5>
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