import java.io.*;

public class Hello {
	public static void main(String[] args) throws Exception {
		BufferedReader bufReader = new BufferedReader(new FileReader("./resources/readme.txt"));
		String line = bufReader.readLine();
		while (line!=null){
			System.out.println(line);
			line = bufReader.readLine();
		}
	}
}
