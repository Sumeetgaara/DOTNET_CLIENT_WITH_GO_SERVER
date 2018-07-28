using System;
using System.Collections.Generic;
using System.Linq;
using System.Net.Http;
using System.Windows.Forms;
using Newtonsoft.Json;
namespace Client
{
	public partial class Form1 : Form
	{
		public Form1()
		{
			InitializeComponent();
		}

		private void button1_Click(object sender, EventArgs e)
		{
			string URI = "http://localhost:8081/blog";
			GetStockInformation(URI);
		}

		public async void GetStockInformation(string url)
		{
			List<blog> stockList = new List<blog>();
			List<blog> dupList = new List<blog>();

			using (var client = new HttpClient())
			{
				using (var response = await client.GetAsync(url))
				{
					if (response.IsSuccessStatusCode)
					{
						var stockJson = await response.Content.ReadAsStringAsync();

						stockList = JsonConvert.DeserializeObject<blog[]>(stockJson).ToList();

						dataGridView1.DataSource = stockList;

					}
				}
			}
		}
	}
}
