import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Typography from '@material-ui/core/Typography';
import Slider from '@material-ui/lab/Slider';

const marks = [
    {
      value: 1,
      label: 'Rare',
    },
    {
      value: 10,
      label: 'Medium',
    },
    {
      value: 17,
      label: 'Well done',
    },
    {
      value: 25,
      label: 'Crispy',
    },
];

function valuetext(value) {
    return `${value}°C`;
}
  
function valueLabelFormat(value) {
return marks.findIndex(mark => mark.value === value) + 1;
}

class FrySlider extends React.Component {
    constructor(props) {
        super(props)
    }


    render() {
        return (
            <div>
                <Typography id="discrete-slider-custom" gutterBottom>
                Custom marks
                </Typography>
                <Slider
                width="300"
                onChangeCommitted={this.props.handleChange}
                getAriaValueText={valuetext}
                defaultValue={15}
                aria-labelledby="discrete-slider-custom"
                step={1}
                valueLabelDisplay="auto"
                marks={marks}
                max={25}
                />
            </div>
        )
    }
}

export default FrySlider;