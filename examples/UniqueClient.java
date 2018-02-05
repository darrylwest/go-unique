package examples;

import java.io.*;
import java.util.*;
import java.net.*;

class UniqueClient {

    public List<String> createCommands() {
        List<String> requests = Arrays.asList( "uuid", "ulid", "guid", "tsid", "txid", "cuid", "xuid", "bytes", "ping", "version" );

        return requests;
    }

    public String sendRequest(PrintWriter out, BufferedReader in, String command) throws Exception {
        out.println(command);

        // read the response
        String response = in.readLine();

        // read the CR and throw it away...
        int ch = in.read();

        return response;
    }

    public static void main(String[] args) throws Exception {
        String host = "localhost";
        int port = 3001;

        UniqueClient client = new UniqueClient();

        try {
            Socket sock = new Socket(host, port);
            PrintWriter out = new PrintWriter(sock.getOutputStream(), true);
            BufferedReader in = new BufferedReader(new InputStreamReader(sock.getInputStream()));

            while (true) {
                for (String cmd : client.createCommands()) {
                    String response = client.sendRequest(out, in, cmd);
                    System.out.println(cmd + "=" + response);
                }

                Thread.sleep(3000);
            }

        } catch(UnknownHostException e) {
            System.err.println("don't know about " + host);
            System.exit(1);
        } catch(IOException e) {
            System.err.println("couldn't get i/o for connection to " + host);
            System.exit(1);
        }
    }
}
