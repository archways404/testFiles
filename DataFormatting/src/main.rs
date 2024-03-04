use std::time::Instant;

#[derive(Debug, Clone)]
struct TableData {
    löneart: String,
    from: String,
    antal: String,
    belopp: String,
}

fn sum_belopp(data: &[TableData]) -> f64 {
    data.iter().map(|item| {
        item.belopp.replace(",", ".").replace(" ", "").parse::<f64>().unwrap_or(0.0)
    }).sum()
}

fn main() {
    let start = Instant::now();

    let table_data = vec![
        TableData { löneart: "0204  Timlön".to_string(), from: "24-03-01".to_string(), antal: "4,00".to_string(), belopp: "888,00".to_string() },
        TableData { löneart: "4200  Sem.ers övriga".to_string(), from: "24-03-01".to_string(), antal: "4,00".to_string(), belopp: "106,56".to_string() },
        // Add more data here based on your requirement
        // The full dataset you provided would be inserted here in the same manner
    ];

    let timlon_data: Vec<_> = table_data.iter().filter(|x| x.löneart == "0204  Timlön").cloned().collect();
    let semers_ovriga_data: Vec<_> = table_data.iter().filter(|x| x.löneart == "4200  Sem.ers övriga").cloned().collect();

    let total_belopp_timlon = sum_belopp(&timlon_data);
    let total_belopp_semers_ovriga = sum_belopp(&semers_ovriga_data);

    let belopp_to_pay = total_belopp_timlon + total_belopp_semers_ovriga;
    println!("Total Belopp for 0204 Timlön: {:.2}", total_belopp_timlon);
    println!("Total Belopp for 4200 Sem.ers övriga: {:.2}", total_belopp_semers_ovriga);
    println!("Belopp to Pay: {:.2}", belopp_to_pay);

    // ACCOUNT FOR THEFT (TAXES)
    let tax = 0.7;
    let belopp_to_pay_after_tax = belopp_to_pay * tax;
    println!("Belopp to Pay (after tax): {:.2}", belopp_to_pay_after_tax);

    let duration = start.elapsed();
    println!("Execution time: {:?}", duration);
}
