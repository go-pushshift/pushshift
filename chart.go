package pushshift

import (
	"fmt"
	"image"
	_ "image/png"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

type ChartQuery struct {
	v       url.Values
	baseURL string
}

// Set scientific notation. (Default: true)
func (q *ChartQuery) ScientificNotation(sci bool) *ChartQuery {
	q.v.Set("chart.set_scientific", strconv.FormatBool(sci))
	return q
}

// Set the color map of the visualization.
// More information about possible color maps: https://matplotlib.org/examples/color/colormaps_reference.html
func (q *ChartQuery) ColorMap(c ColorMap) *ChartQuery {
	q.v.Set("chart.colormap", c.String())
	return q
}

// Set the chart title.
func (q *ChartQuery) ChartTitle(str string) *ChartQuery {
	q.v.Set("chart.title", str)
	return q
}

// Set the chart subtitle.
func (q *ChartQuery) ChartSubtitle(str string) *ChartQuery {
	q.v.Set("chart.subtitle", str)
	return q
}

// Set the font size of the chart subtitle.
func (q *ChartQuery) ChartSubtitleFontSize(px uint) *ChartQuery {
	q.v.Set("chart.subtitle.fontsize", strconv.FormatUint(uint64(px), 10))
	return q
}

// Sort Y-Axis of chart.
func (q *ChartQuery) Sorted() *ChartQuery {
	q.v.Set("chart.sort", "true")
	return q
}

// Use a logarithmic scale for the X-axis.
func (q *ChartQuery) ChartXLogarithmic(log bool) *ChartQuery {
	q.v.Set("chart.xaxis.log", strconv.FormatBool(log))
	return q
}

// Scale down large numbers on the X-axis.
func (q *ChartQuery) ChartXScale(scale float64) *ChartQuery {
	q.v.Set("chart.xaxis.scale", strconv.FormatFloat(scale, 'f', 2, 64))
	return q
}

// Use a logarithmic scale for the Y-axis.
func (q *ChartQuery) ChartYLogarithmic(log bool) *ChartQuery {
	q.v.Set("chart.yaxis.log", strconv.FormatBool(log))
	return q
}

// Put values in buckets of X size such that
// `value = math.floor(value/x)*x`
func (q *ChartQuery) ChartBucketSize(size float64) *ChartQuery {
	q.v.Set("chart.bucket.size", strconv.FormatFloat(size, 'f', 2, 64))
	return q
}

// Set the minimum value of the X-axis.
func (q *ChartQuery) ChartYMin(min float64) *ChartQuery {
	q.v.Set("chart.ymin", strconv.FormatFloat(min, 'f', 2, 64))
	return q
}

// Set the maximum value of the Y-axis.
func (q *ChartQuery) ChartYMax(min float64) *ChartQuery {
	q.v.Set("chart.ymax", strconv.FormatFloat(min, 'f', 2, 64))
	return q
}

// Set the minimum value of the X-axis.
func (q *ChartQuery) ChartXMin(min float64) *ChartQuery {
	q.v.Set("chart.xmin", strconv.FormatFloat(min, 'f', 2, 64))
	return q
}

// Set the chart X-axis label.
func (q *ChartQuery) ChartXLabel(str string) *ChartQuery {
	q.v.Set("chart.xlabel", str)
	return q
}

// Set the chart Y-axis label.
func (q *ChartQuery) ChartYLabel(str string) *ChartQuery {
	q.v.Set("chart.ylabel", str)
	return q
}

// Trim the last value of a chart.
// When doing aggregations, the last bucket can be misleadingly
// low due to the fact that the time period isn't complete.
// This removes that last value.
func (q *ChartQuery) Trim() *ChartQuery {
	q.v.Set("chart.trim", "true")
	return q
}

func (q *ChartQuery) GetImage() (image.Image, error) {
	res, err := q.execute()
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	img, _, err := image.Decode(res.Body)
	if err != nil {
		return nil, err
	}

	return img, nil
}

func (q *ChartQuery) GetPNGBytes() ([]byte, error) {
	res, err := q.execute()
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	buf, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func (q *ChartQuery) execute() (*http.Response, error) {
	u, err := url.Parse(q.baseURL)
	if err != nil {
		return nil, err
	}

	u.RawQuery = q.v.Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	setHTTPHeaders(&req.Header)

	rateLimit()

	res, err := Client.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		res.Body.Close()
		return nil, fmt.Errorf("HTTP status: %s", res.Status)
	}

	if res.Header.Get("Content-Type") != "image/png" {
		res.Body.Close()
		return nil, fmt.Errorf("unexpected data: %s",
			res.Header.Get("Content-Type"))
	}

	return res, nil
}
