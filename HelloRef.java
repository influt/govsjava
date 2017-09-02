import java.io.BufferedReader;
import java.io.FileReader;
import java.io.IOException;

public class HelloRef {
	private final static String FILE_PATH = "./resources/readme.txt";

	public static void main(String[] args) {
		BufferedReader bufReader;
		String line;
		try{
			bufReader = new BufferedReader(new FileReader(FILE_PATH));
			while ((line=bufReader.readLine())!=null){
				System.out.println(line);
			}
		}catch(IOException e){
			e.printStackTrace();
		}
	}
}
