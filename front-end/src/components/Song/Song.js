import React from 'react';

const Song = ({title, url, thumb}) => {
    return (
        <div className='tc dib br3 pa3 ma2 grow bw2 shadow-5 h4 w5'>
            <img src={thumb} /><br/>
            <label>{title}</label><br/>
            <label>{url}</label>
        </div>
    )
}

export default Song;